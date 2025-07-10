package server

import (
	"faqriansyah/siphter/internal/handler"
	"log"
	"net/http"
)

// Run initializes and starts the HTTP server.
func Run(root string, port string) {
	// Create the handler for serving static files
	handler := handler.Handler(root)

	// Register the handler for all paths
	http.Handle("/", handler)

	log.Printf("Server start at port %v from %v", port, root)
	// ListenAndServe blocks until the server stops or an error occurs.
	http.ListenAndServe(port, nil)
}
