package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func init() {
	var err error
	logger := "postgres://agent:password@localhost/"
	dbName := "bookstore"
	sslMode := "?sslmode=disable"
	sourceName := logger + dbName + sslMode
	db, err = sql.Open("postgres", sourceName)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connnection to database established")
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	toQuery := "SELECT * FROM books"
	rows, err := db.Query(toQuery)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	books := make([]Book, 0)

	for rows.Next() {
		buk := Book{}
		err := rows.Scan(&buk.isbn, &buk.title, &buk.author, &buk.price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		books = append(books, buk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, buk := range books {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f \n", buk.isbn, buk.author, buk.title, buk.price)
	}

}

func main() {
	http.HandleFunc("/", booksIndex)
	port := ":8070"
	fmt.Println("Listening And Serving at port", port)
	http.ListenAndServe(port, nil)
}
