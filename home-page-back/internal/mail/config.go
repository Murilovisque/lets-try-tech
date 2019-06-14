package mail

//"github.com/Murilovisque/lets-try-tech/home-page-back/internal/platform"

import (
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/platform"
)

var config mailConfig

func loadConfig() error {
	const configPath = "/opt/ltt/home-page-back/configs/mail.json"
	config = mailConfig{}
	return platform.LoadConfigFromJSONFile(configPath, &config)
}

type mailConfig struct {
	SmtpServerHost      string
	SmtpServerPort      int
	ContactTeamEmail    string
	ContactTeamPassword string
}
