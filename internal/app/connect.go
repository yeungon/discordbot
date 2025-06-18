package app

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DatabaseConnect(postgres_url string) *sql.DB {
	dbConn, err := sql.Open("postgres", postgres_url)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	return dbConn
}
