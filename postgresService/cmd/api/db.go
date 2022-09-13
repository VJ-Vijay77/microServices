package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func  ConnectPostgres() *Config {
	
	db, err := sqlx.Connect("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connection to Postgres success....")
	}

	return &Config{
		Db: db,
	}
}
