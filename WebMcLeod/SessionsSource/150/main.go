package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	logger := "postgres://agent:password@localhost/"
	dbName := "bookstore"
	sslMod := "?sslmode=disable"

	sourceName := logger + dbName + sslMod
	db, err = sql.Open("postgres", sourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database:", dbName)
}

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	author := r.FormValue("author")
	if author == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	query := "SELECT * FROM books WHERE author = $1"
	rows, err := db.Query(query, author)
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
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)

	port := ":8070"
	fmt.Println("LAS at port", port)
	http.ListenAndServe(port, nil)

}
