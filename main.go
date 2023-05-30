package main

import (
	"fmt"
	"net/http"
)

// HelloServer is a function that handles incoming HTTP requests
func HelloServer(w http.ResponseWriter, r *http.Request) {
	// Extract the path from the request URL and write a greeting message to the response writer
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	// Register the HelloServer function to handle incoming requests to the root URL
	http.HandleFunc("/", HelloServer)

	// Start listening for incoming HTTP requests on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// If an error occurs while starting the server, print the error message to the console
		fmt.Println("Error starting server:", err)
	}
}