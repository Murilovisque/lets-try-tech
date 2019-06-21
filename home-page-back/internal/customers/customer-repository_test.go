package customers

import (
	"database/sql"
	"testing"

	"gotest.tools/assert"
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

	assert.NilError(t, AddCustomerMessage(&c), "add customer message failed")
}

func getOldestCustomerMessageShouldReturnsOneRecord(t *testing.T) {
	cs, err := OldestCustomerMessages()
	assert.NilError(t, err, "get oldest customer message failed")
	assert.Assert(t, cs != nil, "oldest customer message did not found")
	assert.Equal(t, 1, len(cs), "There should be only one element")
	assert.Equal(t, uint(1), cs[0].ID, "ID does not match")
	assert.Equal(t, "Teste", cs[0].Name, "Name does not match")
	assert.Equal(t, uint(123), cs[0].Tel, "Tel does not match")
	assert.Equal(t, "teste@domain.com", cs[0].Email, "Email does not match")
	assert.Equal(t, "I have a job for you", cs[0].Message, "Message does not match")
}

func removeCustomerMessageShouldWorks(t *testing.T) {
	c := CustomerMessage{
		ID:      1,
		Name:    "Teste",
		Tel:     123,
		Message: "I have a job for you",
	}
	assert.NilError(t, RemoveCustomerMessage(&c), "remove customer message failed")
}

func getOldestCustomerMessageShouldNotReturnsRecord(t *testing.T) {
	cs, err := OldestCustomerMessages()
	assert.NilError(t, err, "get oldest customer message failed")
	assert.Assert(t, cs == nil, "oldest customer message found")
}

func setupMemoryDbAndTable(t *testing.T) {
	var err error
	db, err = sql.Open("sqlite3", ":memory:")
	assert.NilError(t, err, "Error loading in-memory database")
	assert.NilError(t, db.Ping(), "Error pinging in-memory database")
	assert.NilError(t, createTable(), "Error to create table")
}
