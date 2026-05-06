package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
)

// Boy names list
var boyNames = []string{
	"Liam", "Noah", "Oliver", "Elijah", "William", "James", "Benjamin", "Lucas", "Henry", "Theodore",
	"Alexander", "Michael", "Daniel", "Matthew", "Sebastian", "Jack", "Jayden", "John", "David", "Samuel",
}

// Girl names list
var girlNames = []string{
	"Olivia", "Emma", "Charlotte", "Amelia", "Sophia", "Isabella", "Ava", "Mia", "Ella", "Luna",
	"Camila", "Harper", "Evelyn", "Abigail", "Emily", "Elizabeth", "Sofia", "Mila", "Samantha", "Layla",
}

var randomIntn = secureRandomIntn

func secureRandomIntn(n int) int {
	value, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0
	}

	return int(value.Int64())
}

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

func newServer(publicDir string) http.Handler {
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

	return mux
}

func nameGenerator(w http.ResponseWriter, r *http.Request) {
	requestedGender := strings.TrimSpace(r.URL.Query().Get("gender"))

	// Determine gender based on query parameter
	// When empty, picks randomly from either boy or girl names
	var names []string
	selectedGender := ""

	if requestedGender == "boy" {
		names = boyNames
		selectedGender = "boy"
	} else if requestedGender == "girl" {
		names = girlNames
		selectedGender = "girl"
	} else {
		// Default: pick randomly from all names, then derive the gender from the selected index.
		nameIndex := randomIntn(len(boyNames) + len(girlNames))
		if nameIndex < len(boyNames) {
			names = boyNames
			selectedGender = "boy"
			writeNameResponse(w, requestedGender, selectedGender, names[nameIndex], len(names))
			return
		} else {
			names = girlNames
			selectedGender = "girl"
			writeNameResponse(w, requestedGender, selectedGender, names[nameIndex-len(boyNames)], len(names))
			return
		}
	}

	if len(names) == 0 {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	// Select random name from chosen list
	name := names[randomIntn(len(names))]
	writeNameResponse(w, requestedGender, selectedGender, name, len(names))
}

func writeNameResponse(w http.ResponseWriter, requestedGender, selectedGender, name string, total int) {
	response := map[string]interface{}{
		"name":    name,
		"gender":  selectedGender,
		"total":   total,
		"message": getResponseMessage(requestedGender, selectedGender, name),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getResponseMessage(requestedGender, actualGender, name string) string {
	if requestedGender == "" {
		// Random call with no parameter - return actual gender
		return fmt.Sprintf("Random %s name", actualGender)
	}

	if requestedGender == actualGender {
		return fmt.Sprintf("Requested %s and received %s", requestedGender, requestedGender)
	}

	// Requested one gender but got another
	return fmt.Sprintf("Random %s name (requested: %s)", actualGender, requestedGender)
}
