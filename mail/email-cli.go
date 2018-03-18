package mail

import (
	"github.com/my0sot1s/social/utils"
	gomail "gopkg.in/gomail.v2"
)

type EmailMgr struct {
	MailDial *gomail.Dialer
	HOST     string
	PORT     int
	Username string
	Password string
	// ("smtp.gmail.com", 587, "username", "password")
	// ("smtp.zoho.com", 587, "username", "password")
}

func (m *EmailMgr) Config(host, username, password string, port int) {
	m.HOST = host
	m.Password = password
	m.Username = username
	m.PORT = port
	m.MailDial = gomail.NewDialer(m.HOST, m.PORT, m.Username, m.Password)
	utils.Log("ಠ‿ಠ Gomail ad running ಠ‿ಠ")
}

func (ml *EmailMgr) SendMail(from, body, subject, to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if err := ml.MailDial.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
