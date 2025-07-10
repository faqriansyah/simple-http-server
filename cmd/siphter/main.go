package main

import (
	"faqriansyah/siphter/internal/config"
	"faqriansyah/siphter/internal/server"
	"log"
)

func main() {
	// Parse configuration (root path and port)
	cfg, err := config.ParseArgs()
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	server.Run(cfg.RootPath, cfg.Port)
}
