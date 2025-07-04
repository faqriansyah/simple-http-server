package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config holds the application's configuration parameters.
type Config struct {
	RootPath string
	Port     string
}

// ParseArgs parses command-line arguments to determine the root path and port.
func ParseArgs() (*Config, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %w", err)
	}

	cfg := &Config{
		RootPath: currentPath,
		Port:     ":8080", // Default port
	}

	switch len(os.Args) {
	case 3: // If both PATH and PORT are available
		cfg.RootPath = filepath.Join(currentPath, os.Args[1])
		cfg.Port = fmt.Sprintf(":%s", os.Args[2])
	case 2: // If only PATH is available
		cfg.RootPath = filepath.Join(currentPath, os.Args[1])
	default: // Handle unexpected number of arguments (e.g., more than 3)
		// It's good to inform the user about incorrect usage.
		return nil, fmt.Errorf("usage: %s [path] [port]", os.Args[0])
	}

	return cfg, nil
}
