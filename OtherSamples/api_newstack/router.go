package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates the routes infrastructure to serve our API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range rutas {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)

	}

	return router
}
