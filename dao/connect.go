package dao

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connString string) error {

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(50)

	DB = db
	log.Println("DB connected successfully..")
	return nil
}
