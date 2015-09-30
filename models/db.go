package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// pointer to DB
var db *sql.DB

/* Prepare DB and check if working */
func InitDB(databaseName string) {
	var err error
	db, err := sql.Open("mysql", databaseName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
