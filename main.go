package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func handleArgs() (string, string) {
	// Path where program currently run
	currentPath, _ := os.Getwd()
	path := currentPath
	port := ":8080"

	// Better use cobra??
	switch len(os.Args) {
	case 3: // If both PATH and PORT are available
		path = filepath.Join(currentPath, os.Args[1])
		port = fmt.Sprintf(":%s", os.Args[2])
	case 2: // If only PATH is available
		path = filepath.Join(currentPath, os.Args[1])
	default: // Handle unexpected number of arguments (e.g., more than 3)
		return path, port // Return default port
	}

	return path, port
}

func globalHandler(root string) http.HandlerFunc {
	// Closure for getting a root path
	return func(w http.ResponseWriter, r *http.Request) {
		// Serving file as a response from root folder
		http.ServeFile(w, r, filepath.Clean(filepath.Join(root, r.URL.Path)))
	}
}

func runServer(root string, port string) {
	// The "/" is an entry point for all URL
	http.HandleFunc("/", globalHandler(root))
	http.ListenAndServe(port, nil)
}

func main() {
	root, port := handleArgs()
	fmt.Println("Running on http://localhost:8080")
	runServer(root, port)
}
