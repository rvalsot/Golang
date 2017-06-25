package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Book is a struct that replicates our book registry at db
type Book struct {
	ISBN   string
	Title  string
	Author string
	Price  float32
}

func main() {

	// Connect to db
	source := "postgres://agent:password@localhost/"
	dbName := "bookstore"
	sslMode := "?sslmode=disable"
	dbSource := source + dbName + sslMode
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check if connected
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You are connected to the database: ", dbName)

	// Our query
	toQuery := "SELECT * FROM books;"
	rows, err := db.Query(toQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	books := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		// keep in order the scan!
		err := rows.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			panic(err)
		}
		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, bk := range books {
		fmt.Printf("%s, %s, %s, %.2f \n", bk.ISBN, bk.Title, bk.Author, bk.Price)
	}

}
