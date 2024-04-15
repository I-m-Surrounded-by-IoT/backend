package email

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/emersion/go-sasl"
	smtp "github.com/emersion/go-smtp"
)

func validateSmtpConfig(c *conf.SmtpConfig) error {
	if c == nil {
		return fmt.Errorf("smtp config is nil")
	}
	if c.Host == "" {
		return fmt.Errorf("smtp host is empty")
	}
	if c.Port == 0 {
		return fmt.Errorf("smtp port is empty")
	}
	if c.Protocol == "" {
		return fmt.Errorf("smtp protocol is empty")
	}
	if c.Username == "" {
		return fmt.Errorf("smtp username is empty")
	}
	if c.Password == "" {
		return fmt.Errorf("smtp password is empty")
	}
	if c.From == "" {
		return fmt.Errorf("smtp from is empty")
	}
	return nil
}

func newSmtpClient(c *conf.SmtpConfig) (*smtp.Client, error) {
	var (
		cli *smtp.Client
		err error
	)

	switch strings.ToUpper(c.Protocol) {
	case "TLS":
		cli, err = smtp.DialStartTLS(fmt.Sprintf("%s:%d", c.Host, c.Port), nil)
	case "SSL":
		cli, err = smtp.DialTLS(fmt.Sprintf("%s:%d", c.Host, c.Port), nil)
	default:
		cli, err = smtp.Dial(fmt.Sprintf("%s:%d", c.Host, c.Port))
	}
	if err != nil {
		return nil, fmt.Errorf("dial smtp server failed: %w", err)
	}

	err = cli.Auth(sasl.NewLoginClient(c.Username, c.Password))
	if err != nil {
		cli.Close()
		return nil, fmt.Errorf("auth failed: %w", err)
	}

	return cli, nil
}

type SmtpPool struct {
	mu      sync.RWMutex
	clients []*smtp.Client
	c       *conf.SmtpConfig
	max     int
	active  int
	closed  bool
}

var ErrSmtpPoolClosed = fmt.Errorf("smtp pool closed")

func NewSmtpPool(c *conf.SmtpConfig, max int) (*SmtpPool, error) {
	err := validateSmtpConfig(c)
	if err != nil {
		return nil, err
	}
	return &SmtpPool{
		clients: make([]*smtp.Client, 0, max),
		c:       c,
		max:     max,
	}, nil
}

func (p *SmtpPool) Get() (*smtp.Client, error) {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return nil, ErrSmtpPoolClosed
	}

	if len(p.clients) > 0 {
		cli := p.clients[len(p.clients)-1]
		p.clients = p.clients[:len(p.clients)-1]
		p.active++
		p.mu.Unlock()
		if cli.Noop() != nil {
			cli.Close()
			p.mu.Lock()
			p.active--
			p.mu.Unlock()
			return p.Get()
		}
		return cli, nil
	}

	if p.active >= p.max {
		p.mu.Unlock()
		runtime.Gosched()
		return p.Get()
	}

	cli, err := newSmtpClient(p.c)
	if err != nil {
		p.mu.Unlock()
		return nil, err
	}

	p.active++
	p.mu.Unlock()
	return cli, nil
}

func (p *SmtpPool) Put(cli *smtp.Client) {
	if cli == nil {
		return
	}

	noopErr := cli.Noop()

	p.mu.Lock()
	defer p.mu.Unlock()

	p.active--

	if p.closed || noopErr != nil {
		cli.Close()
		return
	}

	p.clients = append(p.clients, cli)
}

func (p *SmtpPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.closed = true

	for _, cli := range p.clients {
		cli.Close()
	}
	p.clients = nil
}

func (p *SmtpPool) SendEmail(to []string, subject, body string) error {
	cli, err := p.Get()
	if err != nil {
		return err
	}
	defer p.Put(cli)
	return SendEmail(cli, p.getFrom(), to, subject, body)
}

func (p *SmtpPool) getFrom() string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.c.From
}

func (p *SmtpPool) SetFrom(from string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.c.From = from
}
