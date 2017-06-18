package main

import (
	"fmt"
	"io"
	"net/http"
)

func formGet(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
    <form method="get">
      <input type="text" name="q" />
      <input type="submit"/>
    </form>
    <br />`+val)
}

func main() {
	http.HandleFunc("/", formGet)
	http.Handle("/favico.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("To be listed and served at port", port)

	http.ListenAndServe(port, nil)
}
