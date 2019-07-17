package app

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/customers"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/mails"
	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/platform"
	"github.com/pkg/errors"
)

var (
	sendContactUsMessagesChannel      = make(chan customers.CustomerMessage, 5)
	retrySendContactUsMessagesChannel = make(chan customers.CustomerMessage, 5)
	processCounter                    = platform.ProcessCounter{}
	removeCustomerMessage             = customers.RemoveCustomerMessage
	addCustomerMessage                = customers.AddCustomerMessage
	oldestCustomerMessages            = customers.OldestCustomerMessages
	sendMessageToContactTeam          = mails.SendMessageToContactTeam
	setupCustomers                    = customers.Setup
	setupMails                        = mails.Setup
	shutdownCustomers                 = customers.Shutdown
	retrySecondsAfterError            = time.Second * 60 * 5
)

func Setup() error {
	if err := platform.SetupAll(setupCustomers, setupMails); err != nil {
		return err
	}
	defaultProcessError := &ErrApp{http.StatusServiceUnavailable, "Não foi possível processar sua requisição, tente novamente mais tarde"}
	processCounter.Setup(5, defaultProcessError, defaultProcessError)
	startContactUsSavedMessageProcessor()
	startContactUsMessagesSender()
	return nil
}

func startContactUsSavedMessageProcessor() {
	go func() {
		for {
			customers, err := oldestCustomerMessages()
			if err != nil {
				log.Printf("OldestCustomerMessages failed. Retry after %d seconds. Error %v\n", retrySecondsAfterError/time.Second, err)
				time.Sleep(retrySecondsAfterError)
				continue
			}
			for _, c := range customers {
				sendContactUsMessagesChannel <- c
			}
			return
		}
	}()
}

func startContactUsMessagesSender() {
	go func() {
		for c := range sendContactUsMessagesChannel {
			if err := sendMessage(&c); err != nil {
				log.Printf("sendMessage failed. Error %s\n", err)
				retrySendMessage(c)
			}
		}
	}()
}

func retrySendMessage(c customers.CustomerMessage) {
	retrySendContactUsMessagesChannel <- c
	log.Printf("Retry after %d seconds\n", retrySecondsAfterError/time.Second)
	time.AfterFunc(retrySecondsAfterError, func() {
		sendContactUsMessagesChannel <- <-retrySendContactUsMessagesChannel
	})
}

func sendMessage(c *customers.CustomerMessage) error {
	if err := processCounter.IncrementProcess(); err != nil {
		return err
	}
	defer processCounter.DecrementProcess()
	log.Printf("%s's message will be sent\n", c.Name)
	msg := fmt.Sprintf("Contato do cliente: %s\nEmail: %s\nTelefone: %d\nMensagem: %s", c.Name, c.Email, c.Tel, c.Message)
	if err := sendMessageToContactTeam(msg); err != nil {
		return errors.Wrapf(err, "SendMessageToContactTeam failed to send %s\n", msg)
	}
	if err := removeCustomerMessage(c); err != nil {
		return errors.Wrapf(err, "RemoveCustomerMessage failed to remove %v\n", c)
	}
	log.Printf("%s's message sent\n", c.Name)
	return nil
}

func ProcessContactUsMessageReceived(name string, tel uint, email, message string) error {
	if err := processCounter.IncrementProcess(); err != nil {
		return err
	}
	defer processCounter.DecrementProcess()
	if telLength := len(fmt.Sprint(tel)); strings.TrimSpace(name) == "" || telLength < 10 || telLength > 11 || strings.TrimSpace(message) == "" || !regexp.MustCompile(".+@.+").MatchString(email) {
		return &ErrApp{http.StatusBadRequest, "Nome, telefone, email ou mensagem inválidos"}
	}
	c := customers.CustomerMessage{
		Name:    name,
		Tel:     tel,
		Email:   email,
		Message: message,
	}
	if err := addCustomerMessage(&c); err != nil {
		return errors.Wrapf(err, "customers.AddCustomerMessage failed %v", c)
	}
	log.Printf("%s's message added in sender's queue\n", name)
	sendContactUsMessagesChannel <- c
	return nil
}

func Shutdown() {
	log.Println("Finalizing the application...")
	processCounter.Shutdown()
	shutdownCustomers()
	log.Println("Application finalized")
}

type ErrApp struct {
	HTTPStatus int
	msg        string
}

func (e *ErrApp) Error() string {
	return e.msg
}
