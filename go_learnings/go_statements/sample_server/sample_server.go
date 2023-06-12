package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", welcomeHandler) // Register the handler function for the root path
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil) // Start the server
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Local Server!") // Send a welcome message to the client
}
