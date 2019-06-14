package customer

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAndGetOldestAndRemoveCustomerMessageShouldWorks(t *testing.T) {
	setupMemoryDbAndTable(t)
	addCustomerMessageShouldWorks(t)
	getOldestCustomerMessageShouldReturnsOneRecord(t)
	removeCustomerMessageShouldWorks(t)
	getOldestCustomerMessageShouldNotReturnsRecord(t)
}

func addCustomerMessageShouldWorks(t *testing.T) {
	c := CustomerMessage{
		Name:    "Teste",
		Tel:     123,
		Email:   "teste@domain.com",
		Message: "I have a job for you",
	}
	assert.Nil(t, AddCustomerMessage(&c), "add customer message failed")
}

func getOldestCustomerMessageShouldReturnsOneRecord(t *testing.T) {
	c, err := OldestCustomerMessage()
	assert.Nil(t, err, "get oldest customer message failed")
	assert.NotNil(t, c, "oldest customer message did not found")
	assert.EqualValues(t, 1, c.ID, "ID matches")
	assert.EqualValues(t, "Teste", c.Name, "Name matches")
	assert.EqualValues(t, 123, c.Tel, "Tel matches")
	assert.EqualValues(t, "teste@domain.com", c.Email, "Email matches")
	assert.EqualValues(t, "I have a job for you", c.Message, "Message matches")
}

func removeCustomerMessageShouldWorks(t *testing.T) {
	c := CustomerMessage{
		ID:      1,
		Name:    "Teste",
		Tel:     123,
		Message: "I have a job for you",
	}
	assert.Nil(t, RemoveCustomerMessage(&c), "remove customer message failed")
}

func getOldestCustomerMessageShouldNotReturnsRecord(t *testing.T) {
	c, err := OldestCustomerMessage()
	assert.Nil(t, err, "get oldest customer message failed")
	assert.Nil(t, c, "oldest customer message found")
}

func setupMemoryDbAndTable(t *testing.T) {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err, "Error loading in-memory database")
	assert.Nil(t, db.Ping(), "Error pinging in-memory database")
	assert.Nil(t, createTable(), "Error to create table")
}
