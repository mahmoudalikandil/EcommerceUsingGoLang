package dataaccess

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	port     = 5432
	host     = "localhost"
	user     = "postgres"
	password = "Postgres55$"
	dbname   = "postgres"
)

func connectToDatabase() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	db.SetMaxOpenConns(10)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// defer db.Close()
}
func init() {
	connectToDatabase()
}
