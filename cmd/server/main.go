package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yeboahd24/recaptcha-demo/internal/config"
	"github.com/yeboahd24/recaptcha-demo/internal/handlers"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize handlers
	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux, cfg)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
