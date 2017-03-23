package main

import (
	"log"
	"net/http"
	"time"
)

// Logger creates a web log for the actions of our server
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		log.Printf("New log: \t %s \t %s \t %s \t %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
