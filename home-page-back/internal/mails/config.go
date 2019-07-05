package mails

import (
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/platform"
)

var config mailConfig

func loadConfig() error {
	const configPath = "/etc/home-page-back/mail.json"
	config = mailConfig{}
	return platform.LoadConfigFromJSONFile(configPath, &config)
}

type mailConfig struct {
	SmtpServerHost      string
	SmtpServerPort      int
	ContactTeamEmail    string
	ContactTeamPassword string
}
