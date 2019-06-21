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
	chProcessContactUsMessage chan customers.CustomerMessage = make(chan customers.CustomerMessage, 5)
	processCounter            platform.ProcessCounter        = platform.ProcessCounter{}
	shutdownMode                                             = true
	removeCustomerMessage                                    = customers.RemoveCustomerMessage
	addCustomerMessage                                       = customers.AddCustomerMessage
	oldestCustomerMessages                                   = customers.OldestCustomerMessages
	sendMessageToContactTeam                                 = mails.SendMessageToContactTeam
	setupCustomers                                           = customers.Setup
	setupMails                                               = mails.Setup
	shutdownCustomers                                        = customers.Shutdown
)

func Setup() error {
	platform.SetupAll(setupCustomers, setupMails)
	shutdownMode = false
	go processContactUsSavedMessage()
	go sendContactUsMessages()
	return nil
}

func processContactUsSavedMessage() {
	for {
		if shutdownMode {
			break
		}
		customers, err := oldestCustomerMessages()
		if err != nil {
			log.Printf("OldestCustomerMessages failed. It is going to process contact us message again after %d seconds. Error %v\n", 10, err)
			time.Sleep(time.Second * 10)
			continue
		}
		for _, c := range customers {
			chProcessContactUsMessage <- c
		}
		break
	}
}

func sendContactUsMessages() {
	for c := range chProcessContactUsMessage {
		sendMessage(&c)
	}
}

func sendMessage(c *customers.CustomerMessage) {
	processCounter.IncrementProcess()
	defer processCounter.DecrementProcess()
	log.Printf("%s's message will be sent\n", c.Name)
	msg := fmt.Sprintf("Contato do cliente: %s\nEmail: %s\nTelefone: %d\nMensagem: %s", c.Name, c.Email, c.Tel, c.Message)
	if err := sendMessageToContactTeam(msg); err != nil {
		log.Printf("SendMessageToContactTeam failed to send %s. Error %v\n", msg, err)
		return
	}
	if err := removeCustomerMessage(c); err != nil {
		log.Printf("RemoveCustomerMessage failed to remove %v. Error %v\n", c, err)
		return
	}
	log.Printf("%s's message sent\n", c.Name)
}

func ProcessContactUsMessageReceived(name string, tel uint, email, message string) error {
	if shutdownMode {
		return &ErrApp{http.StatusServiceUnavailable, "Não foi possível processar sua requisição, tente novamente mais tarde"}
	}
	processCounter.IncrementProcess()
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
	chProcessContactUsMessage <- c
	return nil
}

func Shutdown() {
	shutdownMode = true
	log.Println("Finalizing the application...")
	processCounter.WaitForAllProcessesToComplete()
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
