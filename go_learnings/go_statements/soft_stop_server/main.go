package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", welcomeHandler)

	go func() {
		fmt.Println("Server listening on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")

	// Attempt to shut down the server gracefully
	if err := server.Shutdown(nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server stopped.")
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Local Server!") // Send a welcome message to the client
}
