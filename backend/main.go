package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var boyNames = []string{
	"Liam", "Noah", "Oliver", "Elijah", "William", "James", "Benjamin", "Lucas", "Henry", "Theodore",
	"Alexander", "Michael", "Daniel", "Matthew", "Sebastian", "Jack", "Jayden", "John", "David", "Samuel",
}

var girlNames = []string{
	"Olivia", "Emma", "Charlotte", "Amelia", "Sophia", "Isabella", "Ava", "Mia", "Ella", "Luna",
	"Camila", "Harper", "Evelyn", "Abigail", "Emily", "Elizabeth", "Sofia", "Mila", "Samantha", "Layla",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	publicDir := "frontend/www"

	mux := http.NewServeMux()

	// Health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
	})

	// Name generator API
	mux.HandleFunc("/api/v1/name-generator", nameGenerator)

	// Static files - serve www/ directory
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(publicDir))))

	fmt.Printf("Starting server on port %s\n", port)
	fmt.Printf("Serving content from: %s\n", publicDir)
	fmt.Printf("Visit: http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func nameGenerator(w http.ResponseWriter, r *http.Request) {
	gender := r.URL.Query().Get("gender")

	var names []string
	selectGender := gender

	if selectGender == "boy" {
		names = boyNames
	} else if selectGender == "girl" {
		names = girlNames
	} else {
		// Default: random gender
		if rand.Intn(2) == 0 {
			names = boyNames
		} else {
			names = girlNames
		}
	}

	if len(names) == 0 {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	idx := rand.Intn(len(names))
	name := names[idx]

	if selectGender == "" {
		selectGender = "boy" // Will be set based on actual name
	}

	response := map[string]interface{}{
		"name":    name,
		"gender":  selectGender,
		"total":   len(names),
		"message": getRandomMessage(selectGender),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getRandomMessage(gender string) string {
	if r := rand.Intn(2); r == 0 {
		return "Random name"
	}
	return fmt.Sprintf("Random %s name", gender)
}
