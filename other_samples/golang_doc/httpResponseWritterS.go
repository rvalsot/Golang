package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Trailer", "AtEnd1, AtEnd2")
		w.Header().Add("Trailer", "AtEnd3")

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		w.Header().Set("AtEnd1", "value 1")
		io.WriteString(w, "This HTTP Response has both headers before this text and trailers at the end. \n")
		w.Header().Set("AtEnd2", "value 2")
		w.Header().Set("AtEnd3", "value 3")
	})
}
