package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func lector(w http.ResponseWriter, r *http.Request) {
	var st string

	fmt.Println(r.Method)
	if r.Method == http.MethodPost {

		// open
		fil, heade, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer fil.Close()

		// to inform:
		fmt.Println("\n File:", fil, "\n Header:", heade, "\n Err:", err)

		// read
		bs, err := ioutil.ReadAll(fil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		st = string(bs)

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
    <form method="post" enctype="multipart/form-data">
      <input type="file" name="q"/>
      <input type="submit"/>
    </form>
    <br /> `+st)
}

func main() {
	http.HandleFunc("/", lector)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listen And Serve", port)
	http.ListenAndServe(port, nil)
}
