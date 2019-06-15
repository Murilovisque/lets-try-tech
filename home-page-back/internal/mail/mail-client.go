package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

// Setup mail to works
func Setup() error {
	err := loadConfig()
	if err == nil {
		log.Println("Mail service set up!")
	}
	return err
}

// SendMessageToContactTeam send a mail message to contact team
func SendMessageToContactTeam(msg string) error {
	auth := smtp.PlainAuth("", config.ContactTeamEmail, config.ContactTeamPassword, config.SmtpServerHost)
	return smtp.SendMail(fmt.Sprintf("%s:%d", config.SmtpServerHost, config.SmtpServerPort), auth,
		config.ContactTeamEmail, []string{config.ContactTeamEmail}, []byte(msg))
}
