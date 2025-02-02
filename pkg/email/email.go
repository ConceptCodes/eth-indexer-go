package email

import (
	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/rs/zerolog"
	"gopkg.in/gomail.v2"
)

type EmailClient struct {
	cfg *config.Config
	log *zerolog.Logger
}

func NewEmailClient(cfg *config.Config, log *zerolog.Logger) *EmailClient {
	return &EmailClient{
		cfg: cfg,
		log: log,
	}
}

func (e *EmailClient) SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.cfg.SmtpFrom)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.cfg.SmtpHost, e.cfg.SmtpPort, e.cfg.SmtpUsername, e.cfg.SmtpPassword)

	if err := d.DialAndSend(m); err != nil {
		e.log.Error().Err(err).Msg("Failed to send email")
		return err
	}

	e.log.Info().Msg("Email sent successfully")
	return nil
}
