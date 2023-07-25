package email

import "log"

type GmailSender struct {
	log *log.Logger
}

func NewGmailSender(l *log.Logger) *GmailSender {
	return &GmailSender{
		log: l,
	}
}

func (g *GmailSender) SendEmail() error {
	return nil
}
