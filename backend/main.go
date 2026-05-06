package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Random boy and girl names
var boyNames = []string{
	"Liam", "Noah", "Oliver", "Elijah", "William", "James", "Benjamin", "Lucas", "Henry", "Theodore",
	"Alexander", "Michael", "Daniel", "Matthew", "Sebastian", "Jack", "Jayden", "John", "David", "Samuel",
}

var girlNames = []string{
	"Olivia", "Emma", "Charlotte", "Amelia", "Sophia", "Isabella", "Ava", "Mia", "Ella", "Luna",
	"Camila", "Harper", "Evelyn", "Abigail", "Emily", "Elizabeth", "Sofia", "Mila", "Samantha", "Layla",
}

// Name API endpoint
func nameGenerator(w http.ResponseWriter, r *http.Request) {
	gender := strings.TrimSpace(r.URL.Query().Get("gender"))
	
	if gender == "" {
		gender = "" // Default: pick any
	}
	
	switch {
	case gender == "boy":
		idx := rand.Intn(len(boyNames))
		response := map[string]interface{}{
			"name":  boyNames[idx],
			"gender": "boy",
			"total": len(boyNames),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		
	case gender == "girl":
		idx := rand.Intn(len(girlNames))
		response := map[string]interface{}{
			"name":  girlNames[idx],
			"gender": "girl",
			"total": len(girlNames),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	
	default:
		// Default: random boy or girl
		if rand.Intn(2) == 0 {
			idx := rand.Intn(len(boyNames))
			response := map[string]interface{}{
				"name":  boyNames[idx],
				"gender": "boy",
				"total": len(boyNames) + len(girlNames),
			}
		} else {
			idx := rand.Intn(len(girlNames))
			response := map[string]interface{}{
				"name":  girlNames[idx],
				"gender": "girl",
				"total": len(boyNames) + len(girlNames),
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve from frontend/www (relative to the backend folder)
	publicDir := "frontend/www"
	
	// Create a custom mux
	mux := http.NewServeMux()
	
	// Register static files for the public directory
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(publicDir))))
	
	// Register health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
	})
	
	// Register name generator API
	mux.HandleFunc("/api/v1/name-generator", nameGenerator)
	
	fmt.Printf("Starting server on port %s\n", port)
	fmt.Printf("Serving content from: %s\n", publicDir)
	fmt.Printf("Visit: http://localhost:%s\n", port)
	fmt.Printf("API: http://localhost:%s/api/v1/name-generator\n", port)
	
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
