package app

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yeungon/discordbot/internal/config"
)

func DatabaseConnect() *sql.DB {
	dbConn, err := sql.Open("postgres", config.PostgreSql_URL())
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	//fmt.Println(dbConn)
	return dbConn
}
