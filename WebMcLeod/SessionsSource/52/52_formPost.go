package main

import (
	"io"
	"net/http"
)

func deployer(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("q")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
    <form method="post">
      <input type="text" name="q" />
      <input type="submit" />
    </form>
    <br />`+val)
}

func main() {
	http.HandleFunc("/", deployer)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	http.ListenAndServe(port, nil)
}
