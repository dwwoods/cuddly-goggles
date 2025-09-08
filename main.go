package main

import (
	"fmt"
	"net/http"
	"encoding/json" // Make sure to import the "encoding/json" package
)
// Task defines the structure of our task items.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

// tasks is our in-memory "database". We'll pre-populate it with some data.
var tasks = []Task{
	{ID: 1, Title: "Learn Go fundamentals", IsCompleted: true},
	{ID: 2, Title: "Build a basic API", IsCompleted: false},
	{ID: 3, Title: "Add a database", IsCompleted: false},
}
// --- NEW HANDLER FUNCTION ---
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Tell the client that we are sending JSON data
	w.Header().Set("Content-Type", "application/json")

	// Encode our tasks slice into JSON and write it to the response writer
	json.NewEncoder(w).Encode(tasks)

}

func main() {
	// Tell Go to handle all requests to the homepage ("/") with the homePage function.
	http.HandleFunc("/tasks", tasksHandler)

	// Start the server on port 8080.
	// The server will run forever until we stop it.
	fmt.Println("Server is starting on port 8080...")
	http.ListenAndServe(":8080", nil)

fmt.Println("\n--- End of Program ---") 
}
