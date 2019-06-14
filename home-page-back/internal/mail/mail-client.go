package mail

import (
	"fmt"
	"net/smtp"
)

// Setup mail to works
func Setup() error {
	return loadConfig()
}

// SendMessageToContactTeam send a mail message to contact team
func SendMessageToContactTeam(msg string) error {
	auth := smtp.PlainAuth("", config.ContactTeamEmail, config.ContactTeamPassword, config.SmtpServerHost)
	return smtp.SendMail(fmt.Sprintf("%s:%d", config.SmtpServerHost, config.SmtpServerPort), auth,
		config.ContactTeamEmail, []string{config.ContactTeamEmail}, []byte(msg))
}
