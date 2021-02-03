package mailer

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	"gopkg.in/mail.v2"
)

type MailConfigs struct {
	Address    string
	Port       int
	Username   string
	Password   string
	FromHeader string
}

type Mail struct {
	Subject string
	To      []string
	CC      []string
	Body    string
}

var (
	mailerInstance *mail.Message
	dialer         *mail.Dialer
)

// Init
func Init(config MailConfigs) {
	m := mail.NewMessage()
	m.SetHeader("From", config.FromHeader)

	mailerInstance = m

	dialer = mail.NewDialer(config.Address, config.Port, config.Username, config.Password)
	fmt.Println(aurora.Green("*** Mailer init successfully"))
}

// SendMail ...
func SendMail(email Mail) (err error) {
	mailerInstance.SetHeader("To", email.To...)
	mailerInstance.SetHeader("Subject", email.Subject)
	mailerInstance.SetBody("text/html", email.Body)
	err = dialer.DialAndSend(mailerInstance)
	if err != nil {
		fmt.Println("Send mail err", err)
	}
	return
}
