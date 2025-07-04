package handler

import (
	"net/http"
)

// StaticFileServer creates an http.HandlerFunc that serves files from the specified root.
func StaticFileServer(root string) http.HandlerFunc {
	// Closure for getting a root path
	return func(w http.ResponseWriter, r *http.Request) {
		// Serving file as a response from root folder,
		// Use http.FileServer for robust static file serving
		// http.FileServer internally handle the request path

		http.FileServer(http.Dir(root)).ServeHTTP(w, r)
	}
}
