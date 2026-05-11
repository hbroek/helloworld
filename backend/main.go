package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	publicDir := "frontend/www"
	mux := newServer(publicDir)

	fmt.Printf("Starting server on port %s\n", port)
	fmt.Printf("Serving content from: %s\n", publicDir)
	fmt.Printf("Visit: http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

// Helper function for CORS enabled status
func hasCorsEnabled() bool {
	if corsEnabled := os.Getenv("CORS_ENABLED"); corsEnabled != "" {
		return corsEnabled == "true"
	}
	return true // Default to enabled
}
