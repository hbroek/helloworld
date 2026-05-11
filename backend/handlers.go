package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Handlers struct {
	names *NameService
}

func NewHandlers(names *NameService) *Handlers {
	return &Handlers{names: names}
}

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func (h *Handlers) NameGenerator(w http.ResponseWriter, r *http.Request) {
	requestedGender := strings.TrimSpace(r.URL.Query().Get("gender"))

	response, err := h.names.Generate(requestedGender)
	if err != nil {
		http.Error(w, "name generator unavailable", http.StatusBadGateway)
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(value)
}
