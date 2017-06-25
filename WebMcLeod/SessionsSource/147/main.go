package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	// This is a basic connection to a DB created in annexed instructions.sql

	dataSource := "postgres://agent:password@localhost/bookstore?sslmode=disable"
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You're connected to the database")
}
