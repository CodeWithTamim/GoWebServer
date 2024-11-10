package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create a file server to serve static files from the "./static" directory
	fileServer := http.FileServer(http.Dir("./static"))

	// Set the file server as the handler for the root URL path
	http.Handle("/", fileServer)

	// Print a message indicating the server is starting
	fmt.Println("Starting server on port :8080")

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // Log any errors that occur while starting the server
	}
}
