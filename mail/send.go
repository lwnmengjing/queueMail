package mail

import (
	"github.com/lwnmengjing/queueMail/form"
	"gopkg.in/gomail.v2"
	"log"
	"net/smtp"
)

func SendMail(message form.MessageMail) (err error) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", message.From.Email, message.From.Name)
	m.SetHeader("To", message.To...)
	m.SetHeader("Subject", message.Subject)
	for _, attachment := range message.Attachments {
		m.Attach(attachment)
	}
	m.SetBody("text/html", message.Body)

	d := gomail.NewDialer(message.Host, message.Port, message.Username, message.Password)
	if message.Auth {
		d.Auth = LoginAuth(message.Username, message.Password)
	}
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
	return
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}
