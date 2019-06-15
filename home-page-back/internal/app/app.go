package app

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/customer"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/mail"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/platform"
	"github.com/pkg/errors"
)

var (
	chProcessContactUsMessage chan *customer.CustomerMessage
	processCounter            platform.ProcessCounter
	shutdownSignalReceived    bool
)

func Setup() error {
	var setups []func() error
	setups = append(setups, func() error { return customer.Setup() })
	setups = append(setups, func() error { return mail.Setup() })
	for _, s := range setups {
		if err := s(); err != nil {
			return err
		}
	}
	chProcessContactUsMessage = make(chan *customer.CustomerMessage, 5)
	processCounter = platform.ProcessCounter{}
	shutdownSignalReceived = false
	go processContactUsSavedMessage()
	go processContactUsMessage()
	return nil
}

func processContactUsSavedMessage() {
	for {
		if shutdownSignalReceived {
			break
		}
		c, err := customer.OldestCustomerMessage()
		if err != nil {
			log.Printf("OldestCustomerMessage failed. It is going to process contact us message again after %d seconds. Error %v\n", 10, err)
			time.Sleep(time.Second * 10)
			continue
		}
		if c == nil {
			break
		}
		chProcessContactUsMessage <- c
	}
}

func processContactUsMessage() {
	for c := range chProcessContactUsMessage {
		if shutdownSignalReceived {
			break
		}
		processCounter.IncrementProcess()
		msg := fmt.Sprintf("Contato do cliente: %s\nEmail: %s\nTelefone: %d\nMensagem: %s", c.Name, c.Email, c.Tel, c.Message)
		if err := mail.SendMessageToContactTeam(msg); err != nil {
			log.Printf("SendMessageToContactTeam failed to send %s. Error %v\n", msg, err)
			processCounter.DecrementProcess()
			continue
		}
		if err := customer.RemoveCustomerMessage(c); err != nil {
			log.Printf("RemoveCustomerMessage failed to remove %v. Error %v\n", c, err)
		}
		processCounter.DecrementProcess()
	}
}

func ContactUsMessageReceived(name string, tel uint, email, message string) error {
	if shutdownSignalReceived {
		return &ErrApp{http.StatusServiceUnavailable, "Não foi possível processar sua requirição, tente novamente mais tarde"}
	}
	processCounter.IncrementProcess()
	defer processCounter.DecrementProcess()
	if telLength := len(fmt.Sprint(tel)); strings.TrimSpace(name) == "" || telLength < 10 || telLength > 11 || strings.TrimSpace(message) == "" || !regexp.MustCompile(".+@.+").MatchString(email) {
		return &ErrApp{http.StatusBadRequest, "Nome, telefone, email ou mensagem inválidos"}
	}
	c := customer.CustomerMessage{
		Name:    name,
		Tel:     tel,
		Email:   email,
		Message: message,
	}
	if err := customer.AddCustomerMessage(&c); err != nil {
		return errors.Wrapf(err, "customer.AddCustomerMessage failed %v", c)
	}
	chProcessContactUsMessage <- &c
	return nil
}

func Shutdown() {
	shutdownSignalReceived = true
	log.Println("Finalizing the application...")
	processCounter.WaitForProcessesToComplete()
	customer.Shutdown()
	log.Println("Application finalized")
}

type ErrApp struct {
	HTTPStatus int
	msg        string
}

func (e *ErrApp) Error() string {
	return e.msg
}
