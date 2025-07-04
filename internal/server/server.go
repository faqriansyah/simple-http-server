package server

import (
	"faqriansyah/siphter/internal/handler"
	"log"
	"net/http"
)

// Run initializes and starts the HTTP server.
func Run(root string, port string) error {
	// Create the handler for serving static files
	staticHandler := handler.StaticFileServer(root)

	// Register the handler for all paths
	http.Handle("/", staticHandler) // Use http.Handle with an http.Handler

	log.Printf("Starting static file server on port %s, serving from %s", port, root)
	// ListenAndServe blocks until the server stops or an error occurs.
	return http.ListenAndServe(port, nil) // nil uses the default ServeMux configured by http.Handle
}
