package mails

import (
	"os"
	"strconv"
	"testing"

	"gotest.tools/assert"
)

func TestSendMessageToContactTeamShouldWorks(t *testing.T) {
	loadConfigForTest(t)
	assert.NilError(t, SendMessageToContactTeam("Unit tests"))
}

func loadConfigForTest(t *testing.T) {
	configMap := make(map[string]string)
	for _, v := range []string{"SMTP_SERVER", "SMTP_PORT", "EMAIL", "PASSWORD"} {
		if envVar := os.Getenv(v); envVar == "" {
			t.Fatalf("Env %s not found", envVar)
			return
		}
		configMap[v] = os.Getenv(v)
	}
	smtpPort, err := strconv.Atoi(configMap["SMTP_PORT"])
	if err != nil {
		t.Fatal("SMTP_PORT variable must e number")
	}
	config = mailConfig{
		ContactTeamEmail:    configMap["EMAIL"],
		ContactTeamPassword: configMap["PASSWORD"],
		SmtpServerHost:      configMap["SMTP_SERVER"],
		SmtpServerPort:      smtpPort,
	}
}
