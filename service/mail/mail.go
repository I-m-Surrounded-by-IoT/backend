package mail

import (
	"context"
	"fmt"
	"strings"

	smtp "github.com/emersion/go-smtp"

	mailApi "github.com/I-m-Surrounded-by-IoT/backend/api/mail"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/emersion/go-sasl"
	log "github.com/sirupsen/logrus"
)

type MailService struct {
	smtpConfig *conf.SmtpConfig
	cli        *smtp.Client
	mailApi.UnimplementedMailServer
}

func forceValidateSmtpConfig(c *conf.SmtpConfig) {
	if c == nil {
		log.Fatal("smtp config is nil")
	}
	if c.Host == "" {
		log.Fatal("smtp host is empty")
	}
	if c.Port == 0 {
		log.Fatal("smtp port is 0")
	}
	if c.Protocol == "" {
		log.Fatal("smtp protocol is empty")
	}
	if c.Username == "" {
		log.Fatal("smtp username is empty")
	}
	if c.Password == "" {
		log.Fatal("smtp password is empty")
	}
	if c.From == "" {
		log.Fatal("smtp from is empty")
	}
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
			return nil, fmt.Errorf("start tls failed: %w", err)
		}
	}

	err = cli.Auth(sasl.NewLoginClient(c.Username, c.Password))
	if err != nil {
		return nil, fmt.Errorf("auth failed: %w", err)
	}

	return cli, nil
}

func NewMailService(c *conf.MailConfig) *MailService {
	if c == nil {
		log.Fatal("mail config is nil")
	}
	forceValidateSmtpConfig(c.Smtp)

	cli, err := newSmtpClient(c.Smtp)
	if err != nil {
		log.Fatalf("new smtp client failed: %v", err)
	}

	s := &MailService{
		smtpConfig: c.Smtp,
		cli:        cli,
	}

	log.Info("mail service started")

	return s
}

func (ms *MailService) SendMail(ctx context.Context, req *mailApi.SendMailReq) (*mailApi.Empty, error) {
	return &mailApi.Empty{},
		SendMail(
			ms.cli,
			ms.smtpConfig.From,
			req.To,
			req.Subject,
			req.Body,
		)
}
