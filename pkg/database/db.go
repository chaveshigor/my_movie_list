package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	_        = godotenv.Load("../../.env")
	host     = os.Getenv("DB_HOST")
	port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbName   = os.Getenv("DB_NAME")
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase() Database {
	return Database{}
}

func (d *Database) OpenConnection() {
	connInfo := fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, username, password, dbName,
	)

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSuccessfully connected to db %s", dbName)
	d.Connection = db
}

func (d *Database) CloseConnection() error {
	err := d.Connection.Close()

	if err != nil {
		return err
	}

	return nil
}
