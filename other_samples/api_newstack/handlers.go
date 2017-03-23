package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index ad
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to this API!")
}

// TodoIndex ad
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	// todos := Todos{
	// 	Todo{Name: "Cow"},
	// 	Todo{Name: "Puppy"},
	// 	Todo{Name: "Axolotl"},
	// }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow //
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}

// TodoCreate creates them
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	// Bad reading
	if err != nil {
		panic(err)
	}
	// Close stuff
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Understood but could not process
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(422) // unprocessable entity
		if errr := json.NewEncoder(w).Encode(err); errr != nil {
			panic(errr)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}
