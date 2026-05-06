package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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
	genderParam := strings.TrimSpace(r.URL.Query().Get("gender"))

	// Determine which names list to use
	var names []string

	if genderParam == "boy" {
		names = boyNames
	} else if genderParam == "girl" {
		names = girlNames
	} else {
		// Default: pick from any list
		names = boyNames
	}

	if len(names) == 0 {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	idx := rand.Intn(len(names))
	name := names[idx]

	// Determine the actual gender of the returned name
	actualGender := detectGender(name)

	response := map[string]interface{}{
		"name":    name,
		"gender":  actualGender,
		"total":   len(names),
		"message": getResponseMessage(genderParam, actualGender),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func detectGender(name string) string {
	boyNames := []string{
		"Liam", "Noah", "Oliver", "Elijah", "William", "James", "Benjamin", "Lucas", "Henry", "Theodore",
		"Alexander", "Michael", "Daniel", "Matthew", "Sebastian", "Jack", "Jayden", "John", "David", "Samuel",
	}
	
	for _, b := range boyNames {
		if b == name {
			return "boy"
		}
	}
	
	return "girl" // All others are girl names
}

func getResponseMessage(requestedGender, actualGender string) string {
	if requestedGender == "" {
		return fmt.Sprintf("Random %s name", actualGender)
	}
	if requestedGender == actualGender {
		return fmt.Sprintf("%s name", requestedGender)
	}
	return fmt.Sprintf("Random %s name (requested: %s)", actualGender, requestedGender)
}
