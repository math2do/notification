package email

import (
	"fmt"
	"io/ioutil"
	"net/smtp"

	"github.com/jordan-wright/email"
	log "github.com/sirupsen/logrus"
	"math2do.in/notification/configs"
	"math2do.in/notification/dtos"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error
}

type GmailSender struct {
	log    *log.Entry
	config *configs.EmailConfigs
}

func NewGmailSender() *GmailSender {
	return &GmailSender{
		config: &configs.Config.Email,

		log: log.WithFields(log.Fields{
			"service": "gmail",
		}),
	}
}

func (g *GmailSender) SendEmailForVerification(req dtos.VerifyEmailReq) error {

	data, err := ioutil.ReadFile("/home/math2do/Projects/go/src/math2do.in/notification/static/email.html")
	if err != nil {
		g.log.Println("File reading error", err)
		return err
	}

	return g.SendEmail(
		"math2do: please verify your email",
		string(data),
		[]string{req.To},
		[]string{},
		[]string{},
		[]string{},
	)
}

func (g *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", g.config.Name, g.config.From)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			g.log.Errorf("failed to attach file %s: %w", f, err)
			return err
		}
	}

	smtpAuth := smtp.PlainAuth("", g.config.From, g.config.Password, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
