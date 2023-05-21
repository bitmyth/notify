package email

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"strings"
)

const (
	smtp163Host = "smtp.163.com"
	smtp163Port = "25"
)

type Mail struct {
	From     string
	To       []string
	Password string
	Port     int
	Host     string
}

func NewMail(options ...Option) *Mail {
	m := &Mail{}
	for _, o := range options {
		o(m)
	}
	return m
}

type Option func(m *Mail)

func From(from string) Option {
	return func(m *Mail) {
		m.From = from
	}
}

func To(to []string) Option {
	return func(m *Mail) {
		m.To = to
	}
}

func Password(password string) Option {
	return func(m *Mail) {
		m.Password = password
	}
}

func SmtpHost(Host string) Option {
	return func(m *Mail) {
		m.Host = Host
	}
}

func SmtpPort(port int) Option {
	return func(m *Mail) {
		m.Port = port
	}
}

func (m *Mail) SetFrom(from string) *Mail {
	m.From = from
	return m
}

func (m *Mail) SetTo(to ...string) *Mail {
	m.To = to
	return m
}

func (m *Mail) SetPassword(pass string) *Mail {
	m.Password = pass
	return m
}

func (m *Mail) SetSmtpHost(host string) *Mail {
	m.Host = host
	return m
}

func (m *Mail) SetSmtpPort(port int) *Mail {
	m.Port = port
	return m
}

func (m Mail) Send(messageText string) error {
	// Sender data.
	from := m.From
	password := m.Password

	// Receiver email address.
	to := m.To

	// smtp server configuration.
	smtpHost := m.Host

	smtpPort := m.Port

	// Message.
	//message := []byte(messageText)

	//// Authentication.
	//auth := smtp.PlainAuth("", from, password, smtpHost)
	//
	//// Sending email.
	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	//if err != nil {
	//	return err
	//}

	mm := gomail.NewMessage()

	// Set E-Mail sender
	mm.SetHeader("From", from)

	// Set E-Mail receivers
	mm.SetHeader("To", strings.Join(to, ","))

	// Set E-Mail subject
	mm.SetHeader("Subject", messageText)

	// Set E-Mail body. You can set plain text or html with text/html
	mm.SetBody("text/plain", messageText)

	// Settings for SMTP server
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(mm); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
