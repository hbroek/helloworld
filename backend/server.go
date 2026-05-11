package main

import "net/http"

func newServer(publicDir string) http.Handler {
	nameService := NewNameService(defaultBoyNames, defaultGirlNames, secureRandomIntn)
	handlers := NewHandlers(nameService)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/api/v1/name-generator", handlers.NameGenerator)
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(publicDir))))

	return mux
}
