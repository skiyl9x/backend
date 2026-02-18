package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Task structure for our JSON data
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	// 1. Serve static files (Frontend) from the /static folder
	// Accessing http://localhost:8080/ will look for index.html in ./static
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 2. API Endpoint (Backend)
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks := []Task{
			{ID: 1, Title: "Install Go on RHEL 8", Status: "Done"},
			{ID: 2, Title: "Configure K8s Pods", Status: "In Progress"},
			{ID: 3, Title: "Fix SSH Tunnels", Status: "Todo"},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
