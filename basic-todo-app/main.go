package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`        // Unique ID for each task
	Title     string `json:"title"`     // Title of the task
	Completed bool   `json:"completed"` // Whether the task is done or not
}

var todos []Todo
var nextID = 1

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response type to JSON

	// Parse the input JSON into a Todo struct
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Assign a unique ID and add to the slice
	newTodo.ID = nextID
	nextID++
	todos = append(todos, newTodo)

	// Respond with the newly created task
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response type to JSON

	// Respond with all the todos
	json.NewEncoder(w).Encode(todos)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response type to JSON

	// Parse the input JSON into a Todo struct
	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Find the task by ID and update it
	for i, todo := range todos {
		if todo.ID == updatedTodo.ID {
			todos[i] = updatedTodo
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}

	// If not found, respond with an error
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response type to JSON

	// Parse the input JSON to get the ID of the task to delete
	var id struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Find and delete the task by ID
	for i, todo := range todos {
		if todo.ID == id.ID {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	// If not found, respond with an error
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			createTodoHandler(w, r)
		case "GET":
			getTodosHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/todos/update", updateTodoHandler)
	http.HandleFunc("/todos/delete", deleteTodoHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
