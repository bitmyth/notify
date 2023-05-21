package email

import "testing"

func TestMail_Send(t *testing.T) {
	m := NewMail().
		SetFrom("fishis@163.com").SetTo("fishis@163.com").
		SetSmtpHost("smtp.163.com").SetSmtpPort(25).
		SetPassword("YOURPASSWORD")

	err := m.Send("hello world")
	if err != nil {
		t.Error(err)
		return
	}
}
