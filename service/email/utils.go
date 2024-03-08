package email

import (
	"fmt"
	"strings"

	smtp "github.com/emersion/go-smtp"
)

func FormatMail(from string, to []string, subject, body any) string {
	return fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%v",
		from,
		strings.Join(to, ","),
		subject,
		body,
	)
}

func SendMail(cli *smtp.Client, from string, to []string, subject, body string) error {
	return cli.SendMail(
		from,
		to,
		strings.NewReader(
			FormatMail(
				from,
				to,
				subject,
				body,
			),
		),
	)
}
