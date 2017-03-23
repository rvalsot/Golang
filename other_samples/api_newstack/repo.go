package main

import "fmt"

var currentID int
var todos Todos

// A small seed data

func init() {
	RepoCreateTodo(Todo{Name: "Hedgehog1"})
	RepoCreateTodo(Todo{Name: "Hedgehog2"})
	RepoCreateTodo(Todo{Name: "Hedgehog3"})
	RepoCreateTodo(Todo{Name: "Hedgehog4"})
}

// RepoFindTodo finds them in our list
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	return Todo{}
}

// RepoCreateTodo add them our list
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo removes them from our list
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find todo with ID of %d to delete it", id)
}
