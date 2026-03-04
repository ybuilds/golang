package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open("postgres", "host=localhost user=postgres password=kjm40438 dbname=blog sslmode=disable")
	if err != nil {
		log.Fatalln("error creating db connection")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln("error pinging db")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
