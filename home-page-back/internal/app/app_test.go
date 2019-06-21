package app

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"gotest.tools/assert/cmp"

	"gotest.tools/assert"

	"github.com/Murilovisque/lets-try-tech/home-page-back/internal/customers"
)

const (
	expectedName, expectedTel, expectedEmail, expectedMessage     = "Mu", uint(1332372614), "murilo@domainio.com", "I have a job for you"
	expectedName2, expectedTel2, expectedEmail2, expectedMessage2 = "Ar", uint(1122272612), "arantes@domainio.com", "I have a other job for you"
)

var (
	expectedEmailMessage         = fmt.Sprintf("Contato do cliente: %s\nEmail: %s\nTelefone: %d\nMensagem: %s", expectedName, expectedEmail, expectedTel, expectedMessage)
	expectedEmailMessage2        = fmt.Sprintf("Contato do cliente: %s\nEmail: %s\nTelefone: %d\nMensagem: %s", expectedName2, expectedEmail2, expectedTel2, expectedMessage2)
	expectedCustomerMessage      = customers.CustomerMessage{ID: 1, Name: expectedName, Email: expectedEmail, Tel: expectedTel, Message: expectedMessage}
	expectedCustomerMessage2     = customers.CustomerMessage{ID: 2, Name: expectedName2, Email: expectedEmail2, Tel: expectedTel2, Message: expectedMessage2}
	sendMessageToContactTeamArgs []string
	removeCustomerMessageArgs    []customers.CustomerMessage
	addCustomerMessageArgs       []customers.CustomerMessage
	timeout                      time.Duration
	delay                        time.Duration
)

func init() {
	clearVals()
}

func TestContactUsMessageReceivedShouldWorks(t *testing.T) {
	setupMocks(t)
	err := ProcessContactUsMessageReceived(expectedName, expectedTel, expectedEmail, expectedMessage)
	assert.NilError(t, err)
	waitUntil(t, func() bool { return len(addCustomerMessageArgs) == 1 }, "Expected 1 call to addCustomerMessage - Timeout")
	assert.DeepEqual(t, addCustomerMessageArgs[0], expectedCustomerMessage)
	waitUntil(t, func() bool { return len(sendMessageToContactTeamArgs) == 1 }, "Expected 1 call to sendMessageToContactTeam - Timeout")
	assert.Equal(t, sendMessageToContactTeamArgs[0], expectedEmailMessage)
	waitUntil(t, func() bool { return len(removeCustomerMessageArgs) == 1 }, "Expected 1 call to removeCustomerMessage - Timeout")
	assert.DeepEqual(t, removeCustomerMessageArgs[0], expectedCustomerMessage)
	clearVals()
}

func TestContactUsSavedMessageShouldWorks(t *testing.T) {
	addCustomerMessageArgs = append(addCustomerMessageArgs, expectedCustomerMessage, expectedCustomerMessage2)
	setupMocks(t)
	waitUntil(t, func() bool { return len(sendMessageToContactTeamArgs) == 2 }, "Expected 2 call to sendMessageToContactTeam - Timeout")
	assert.Equal(t, sendMessageToContactTeamArgs[0], expectedEmailMessage)
	assert.Equal(t, sendMessageToContactTeamArgs[1], expectedEmailMessage2)
	waitUntil(t, func() bool { return len(removeCustomerMessageArgs) == 2 }, "Expected 2 call to removeCustomerMessage - Timeout")
	assert.DeepEqual(t, removeCustomerMessageArgs[0], expectedCustomerMessage)
	assert.DeepEqual(t, removeCustomerMessageArgs[1], expectedCustomerMessage2)
	clearVals()
}

func TestShutdownWaitAllProcessFinish(t *testing.T) {
	delay = time.Second
	addCustomerMessageArgs = append(addCustomerMessageArgs, expectedCustomerMessage)
	setupMocks(t)
	waitUntil(t, func() bool { return len(sendMessageToContactTeamArgs) == 1 }, "Expected 1 call to sendMessageToContactTeam - Timeout")
	Shutdown()
	assert.Equal(t, sendMessageToContactTeamArgs[0], expectedEmailMessage)
	waitUntil(t, func() bool { return len(removeCustomerMessageArgs) == 1 }, "Expected 1 call to removeCustomerMessage - Timeout")
	assert.DeepEqual(t, removeCustomerMessageArgs[0], expectedCustomerMessage)
	clearVals()
}

func TestProcessContactUsMessageReceivedShouldReturnBadRequestWhenArgsAreInvalid(t *testing.T) {
	setupMocks(t)
	var err error
	checker := comparisonErrApp(func() error { return err }, http.StatusBadRequest, "Nome, telefone, email ou mensagem inválidos")
	err = ProcessContactUsMessageReceived("", expectedTel, expectedEmail, expectedMessage)
	assert.Assert(t, checker)
	err = ProcessContactUsMessageReceived(expectedName, 1, expectedEmail, expectedMessage)
	assert.Assert(t, checker)
	err = ProcessContactUsMessageReceived(expectedName, expectedTel, "invalidemail", expectedMessage)
	assert.Assert(t, checker)
	err = ProcessContactUsMessageReceived(expectedName, expectedTel, expectedEmail, "")
	assert.Assert(t, checker)
	clearVals()
}

func TestProcessContactUsMessageReceivedShouldReturnUnavailable(t *testing.T) {
	Shutdown()
	var err error
	checker := comparisonErrApp(func() error { return err }, http.StatusServiceUnavailable, "Não foi possível processar sua requisição, tente novamente mais tarde")
	err = ProcessContactUsMessageReceived(expectedName, expectedTel, expectedEmail, expectedMessage)
	assert.Assert(t, checker)
}

func setupMocks(t *testing.T) {
	mockDependenciesFuncs(t)
	mockAndCallSetup(t)
}

func comparisonErrApp(err func() error, status int, msg string) func() cmp.Result {
	return func() cmp.Result {
		errApp, ok := err().(*ErrApp)
		if !ok {
			return cmp.ResultFailure(fmt.Sprintf("Error %v is not ErrApp", err()))
		}
		if errApp.HTTPStatus != status || errApp.msg != msg {
			return cmp.ResultFailure(fmt.Sprintf("Error %v does not have bad request status", err()))
		}
		return cmp.ResultSuccess
	}
}

func mockAndCallSetup(t *testing.T) {
	setupCustomers = func() error { return nil }
	setupMails = func() error { return nil }
	assert.NilError(t, Setup())
}

func mockDependenciesFuncs(t *testing.T) {
	if delay > 0 {
		timeout += delay
	}
	oldestCustomerMessages = func() ([]customers.CustomerMessage, error) { return addCustomerMessageArgs, nil }
	sendMessageToContactTeam = func(msg string) error {
		t.Log("sendMessageToContactTeam called with " + msg)
		sendMessageToContactTeamArgs = append(sendMessageToContactTeamArgs, msg)
		time.Sleep(delay)
		return nil
	}
	removeCustomerMessage = func(c *customers.CustomerMessage) error {
		t.Logf("removeCustomerMessage called with %v", c)
		removeCustomerMessageArgs = append(removeCustomerMessageArgs, *c)
		time.Sleep(delay)
		return nil
	}
	addCustomerMessage = func(c1 *customers.CustomerMessage) error {
		var id uint
		return func(c2 *customers.CustomerMessage) error {
			t.Logf("addCustomerMessage called with %v", c2)
			id++
			c2.ID = id
			addCustomerMessageArgs = append(addCustomerMessageArgs, *c2)
			time.Sleep(delay)
			return nil
		}(c1)
	}
	shutdownCustomers = func() {}
}

func clearVals() {
	sendMessageToContactTeamArgs = []string{}
	removeCustomerMessageArgs = []customers.CustomerMessage{}
	addCustomerMessageArgs = []customers.CustomerMessage{}
	timeout = time.Second
	delay = 0
}

func waitUntil(t *testing.T, condition func() bool, failMsg string) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan struct{}, 1)
	go func() {
		for {
			if condition() {
				ch <- struct{}{}
				break
			}
		}
	}()
	select {
	case <-ctx.Done():
		t.Error(failMsg)
		t.FailNow()
	case <-ch:
	}
}
