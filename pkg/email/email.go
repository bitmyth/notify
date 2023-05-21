package email

import (
	"fmt"
	"net/smtp"
)

const (
	smtp163Host = "smtp.163.com"
	smtp163Port = "25"
)

type Mail struct {
	From     string
	To       []string
	Password string
	Port     string
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

func SmtpPort(port string) Option {
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

func (m *Mail) SetSmtpPort(port string) *Mail {
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
	message := []byte(messageText)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Email Sent Successfully!")

	return nil
}
