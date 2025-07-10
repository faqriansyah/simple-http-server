package handler

import (
	"log"
	"net/http"
)

func serveFile(root string, w http.ResponseWriter, r *http.Request) {
	// Serving file as a response from root folder,
	// Use http.FileServer for robust static file serving
	// http.FileServer internally handle the request path

	http.FileServer(http.Dir(root)).ServeHTTP(w, r)
}

// StaticFileServer creates an http.HandlerFunc that serves files from the specified root.
func Handler(root string) http.HandlerFunc {
	// Closure for getting a root path
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Receiving request")

		serveFile(root, w, r)
	}
}
