package email

import (
	"context"

	emailApi "github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	log "github.com/sirupsen/logrus"
)

type EmailService struct {
	smtpConfig *conf.SmtpConfig
	smtpPool   *SmtpPool
	emailApi.UnimplementedEmailServer
}

func NewEmailService(c *conf.EmailConfig) *EmailService {
	if c == nil {
		log.Fatal("mail config is nil")
	}

	sp, err := NewSmtpPool(c.Smtp, 10)
	if err != nil {
		log.Fatalf("create smtp pool failed: %v", err)
	}

	s := &EmailService{
		smtpConfig: c.Smtp,
		smtpPool:   sp,
	}

	log.Info("email service started")

	return s
}

func (ms *EmailService) SendEmail(ctx context.Context, req *emailApi.SendEmailReq) (*emailApi.Empty, error) {
	cli, err := ms.smtpPool.Get()
	if err != nil {
		return nil, err
	}
	defer ms.smtpPool.Put(cli)
	return &emailApi.Empty{},
		SendMail(
			cli,
			ms.smtpConfig.From,
			req.To,
			req.Subject,
			req.Body,
		)
}
