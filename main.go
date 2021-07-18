package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dev"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping was fine")
}
