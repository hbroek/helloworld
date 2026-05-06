package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	server := newServer("../frontend/www")
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if response.Body.String() != "OK\n" {
		t.Fatalf("expected health body %q, got %q", "OK\n", response.Body.String())
	}
}

func TestNameGeneratorReturnsRequestedBoyName(t *testing.T) {
	restoreRandom := stubRandomIntn(t, 1)
	defer restoreRandom()

	response := requestName(t, "/api/v1/name-generator?gender=boy")

	if response.Name != "Noah" {
		t.Fatalf("expected indexed boy name Noah, got %q", response.Name)
	}
	if response.Gender != "boy" {
		t.Fatalf("expected gender boy, got %q", response.Gender)
	}
	if response.Total != len(boyNames) {
		t.Fatalf("expected total %d, got %d", len(boyNames), response.Total)
	}
}

func TestNameGeneratorReturnsRequestedGirlName(t *testing.T) {
	restoreRandom := stubRandomIntn(t, 2)
	defer restoreRandom()

	response := requestName(t, "/api/v1/name-generator?gender=girl")

	if response.Name != "Charlotte" {
		t.Fatalf("expected indexed girl name Charlotte, got %q", response.Name)
	}
	if response.Gender != "girl" {
		t.Fatalf("expected gender girl, got %q", response.Gender)
	}
	if response.Total != len(girlNames) {
		t.Fatalf("expected total %d, got %d", len(girlNames), response.Total)
	}
}

func TestNameGeneratorRandomGenderUsesSelectedList(t *testing.T) {
	restoreRandom := stubRandomIntn(t, len(boyNames)+3)
	defer restoreRandom()

	response := requestName(t, "/api/v1/name-generator")

	if response.Name != "Amelia" {
		t.Fatalf("expected indexed girl name Amelia, got %q", response.Name)
	}
	if response.Gender != "girl" {
		t.Fatalf("expected gender girl, got %q", response.Gender)
	}
	if response.Total != len(girlNames) {
		t.Fatalf("expected total %d, got %d", len(girlNames), response.Total)
	}
}

func TestNameGeneratorRandomGenderCanReturnBoyName(t *testing.T) {
	restoreRandom := stubRandomIntn(t, 4)
	defer restoreRandom()

	response := requestName(t, "/api/v1/name-generator")

	if response.Name != "William" {
		t.Fatalf("expected indexed boy name William, got %q", response.Name)
	}
	if response.Gender != "boy" {
		t.Fatalf("expected gender boy, got %q", response.Gender)
	}
	if response.Total != len(boyNames) {
		t.Fatalf("expected total %d, got %d", len(boyNames), response.Total)
	}
}

func requestName(t *testing.T, target string) nameResponse {
	t.Helper()

	request := httptest.NewRequest(http.MethodGet, target, nil)
	response := httptest.NewRecorder()

	nameGenerator(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if contentType := response.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected content type application/json, got %q", contentType)
	}

	var payload nameResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	return payload
}

func stubRandomIntn(t *testing.T, values ...int) func() {
	t.Helper()

	original := randomIntn
	call := 0
	randomIntn = func(n int) int {
		if call >= len(values) {
			t.Fatalf("randomIntn called more than %d times", len(values))
		}

		value := values[call]
		call++

		if value < 0 || value >= n {
			t.Fatalf("stubbed random value %d outside range [0,%d)", value, n)
		}

		return value
	}

	return func() {
		randomIntn = original
	}
}

type nameResponse struct {
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Total   int    `json:"total"`
	Message string `json:"message"`
}
