package main

import (
	"faqriansyah/siphter/internal/config"
	"faqriansyah/siphter/internal/server"
	"fmt"
	"log"
)

func main() {
	// Parse configuration (root path and port)
	cfg, err := config.ParseArgs()
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	fmt.Printf("Serving files from: %s\n", cfg.RootPath)
	fmt.Printf("Running on http://localhost%s\n", cfg.Port)

	// Run the server
	if err := server.Run(cfg.RootPath, cfg.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
