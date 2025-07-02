package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// Storing root folder
var root string

func runServer() {
	path, port := handleArgs()
	root = path

	http.HandleFunc("/", globalHandler)
	http.ListenAndServe(port, nil)
}

func handleArgs() (string, string) {
	currentPath, _ := os.Getwd() // HANDLE THE ERROR
	path := currentPath
	port := ":8080"

	switch len(os.Args) {
	case 3: // If both PATH and PORT are available
		path = filepath.Join(currentPath, os.Args[1])
		port = fmt.Sprintf(":%s", os.Args[2])
	case 2: // If only PATH is available
		path = filepath.Join(currentPath, os.Args[1])
	default: // Handle unexpected number of arguments (e.g., more than 3)
		return path, port // Return default port and empty path, or handle as needed
	}

	return path, port
}

func globalHandler(w http.ResponseWriter, r *http.Request) {
	// Serving file as a response from root folder
	http.ServeFile(w, r, filepath.Clean(filepath.Join(root, r.URL.Path)))
}

func main() {
	runServer()
}
