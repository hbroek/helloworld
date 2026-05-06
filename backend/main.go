package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	publicDir := "frontend/www"

	// Handle index.html redirects for directory access
	http.Handle("/", http.StripPrefix("", http.FileServer(http.Dir(publicDir))))

	health := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
	}
	http.Handle("/health", http.HandlerFunc(health))

	fmt.Printf("Starting server on port %s\n", port)
	fmt.Printf("Serving content from: %s\n", publicDir)
	fmt.Printf("Visit: http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}