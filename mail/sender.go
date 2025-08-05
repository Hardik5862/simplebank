package mail

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

const (
	gmailSmtpAuthAddress = "smtp.gmail.com"
)

type EmailSender interface {
	SendMail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendMail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
) error {
	message := mail.NewMsg()

	err := message.From(fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress))
	if err != nil {
		return fmt.Errorf("failed to set from email value: %w", err)
	}

	message.Subject(subject)
	message.SetBodyString(mail.TypeTextHTML, content)

	err = message.To(to...)
	if err != nil {
		return fmt.Errorf("failed to set to email value: %w", err)
	}

	err = message.Cc(cc...)
	if err != nil {
		return fmt.Errorf("failed to set cc email value: %w", err)
	}

	err = message.Bcc(bcc...)
	if err != nil {
		return fmt.Errorf("failed to set bcc email value: %w", err)
	}

	client, err := mail.NewClient(
		gmailSmtpAuthAddress,
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(sender.fromEmailAddress),
		mail.WithPassword(sender.fromEmailPassword),
	)
	if err != nil {
		return fmt.Errorf("failed to create mail client: %w", err)
	}

	err = client.DialAndSend(message)
	if err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}

	return nil
}
