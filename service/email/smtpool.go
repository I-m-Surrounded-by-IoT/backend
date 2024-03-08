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
	cli, err := smtp.Dial(fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return nil, fmt.Errorf("dial smtp server failed: %w", err)
	}

	switch strings.ToUpper(c.Protocol) {
	case "TLS", "SSL":
		err = cli.StartTLS(nil)
		if err != nil {
			cli.Close()
			return nil, fmt.Errorf("start tls failed: %w", err)
		}
	}

	err = cli.Auth(sasl.NewLoginClient(c.Username, c.Password))
	if err != nil {
		cli.Close()
		return nil, fmt.Errorf("auth failed: %w", err)
	}

	return cli, nil
}

type SmtpPool struct {
	mu      sync.Mutex
	clients []*smtp.Client
	c       *conf.SmtpConfig
	max     int
	active  int
}

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

	if len(p.clients) > 0 {
		cli := p.clients[len(p.clients)-1]
		p.clients = p.clients[:len(p.clients)-1]
		if cli.Noop() != nil {
			p.mu.Unlock()
			cli.Close()
			return p.Get()
		}
		p.active++
		p.mu.Unlock()
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
	p.mu.Lock()
	defer p.mu.Unlock()

	p.active--

	if cli == nil {
		return
	}
	if cli.Noop() != nil {
		cli.Close()
		return
	}

	p.clients = append(p.clients, cli)
}
