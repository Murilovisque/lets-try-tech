package customer

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var db *sql.DB

// Setup database
func Setup() error {
	const dbPath = "/opt/ltt/home-page-back/dbs/home-page.db"
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return errors.Wrapf(err, "Error loading database %s", dbPath)
	}
	if err = db.Ping(); err != nil {
		return errors.Wrapf(err, "Error ping database %s", dbPath)
	}
	if err = createTable(); err != nil {
		return errors.Wrap(err, "Error create table customer_message")
	}
	log.Println("Customer service set up!")
	return nil
}

func AddCustomerMessage(c *CustomerMessage) error {
	stmt, err := db.Prepare("INSERT INTO customer_message (name, tel, email, message) values (?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(c.Name, c.Tel, c.Email, c.Message)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = uint(lastID)
	return nil
}

func RemoveCustomerMessage(c *CustomerMessage) error {
	stmt, err := db.Prepare("DELETE FROM customer_message WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.ID)
	return err
}

func OldestCustomerMessage() (*CustomerMessage, error) {
	var c CustomerMessage
	row := db.QueryRow("select id, name, tel, email, message from customer_message order by 1 limit 1")
	err := row.Scan(&c.ID, &c.Name, &c.Tel, &c.Email, &c.Message)
	if err == nil {
		return &c, nil
	} else if err == sql.ErrNoRows {
		return nil, nil
	}
	return nil, err
}

func createTable() error {
	query := `CREATE TABLE IF NOT EXISTS customer_message(
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		tel INTEGER NOT NULL,
		email TEXT NOT NULL,
		message TEXT NOT NULL)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

// Shutdown database
func Shutdown() {
	db.Close()
}

type CustomerMessage struct {
	ID      uint
	Name    string
	Tel     uint
	Email   string
	Message string
}
