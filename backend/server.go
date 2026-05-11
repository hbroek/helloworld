package main

import (
	"net/http"
	"os"
	"strings"
)

func newServer(publicDir string) http.Handler {
	nameService := NewNameService(defaultBoyNames, defaultGirlNames, secureRandomIntn)
	handlers := NewHandlers(nameService)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/api/v1/name-generator", handlers.NameGenerator)
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(publicDir))))

	// Apply CORS headers to all responses
	return withCORS(mux)
}

// withCORS adds CORS headers to all responses
func withCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		
		// Check if CORS is enabled
		if os.Getenv("CORS_ENABLED") == "false" {
			// Still set basic headers for same-origin
			w.Header().Set("Vary", "Origin")
		} else {
			// Set allowed origins from env or default
			allowedOrigins := getAllowedOrigins()
			
			// Check if origin is allowed or if no origin header (same-origin)
			if origin == "" || isOriginAllowed(origin, allowedOrigins) || origin == "*" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")
				w.Header().Set("Vary", "Origin")
			} else {
				// Origin not allowed - return 403
				if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusForbidden)
					return
				}
				http.Error(w, "CORS policy error", http.StatusForbidden)
				return
			}
		}
		
		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.Header().Set("Vary", "Origin")
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Call next handler
		handler.ServeHTTP(w, r)
	})
}

// getAllowedOrigins returns the list of allowed origins
func getAllowedOrigins() []string {
	envOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if envOrigins != "" {
		origins := strings.Split(envOrigins, ",")
		result := make([]string, len(origins))
		for _, o := range origins {
			trimmed := strings.TrimSpace(o)
			if trimmed != "" {
				result = append(result, trimmed)
			}
		}
		if len(result) > 0 {
			return result
		}
	}
	
	return defaultCorsAllowedOrigins
}

// defaultCorsAllowedOrigins - default allowed origins
var defaultCorsAllowedOrigins = []string{
	"http://localhost:*",
	"http://127.0.0.1:*",
	"https://*.onrender.com",
	"https://*.vercel.app",
}

// isOriginAllowed checks if the origin matches any in the allowed list
func isOriginAllowed(origin string, allowed []string) bool {
	if origin == "" {
		return true
	}
	
	for _, allowedOrigin := range allowed {
		// Handle wildcard ports
		if strings.Contains(origin, allowedOrigin) {
			return true
		}
		// Handle wildcard (*)
		if allowedOrigin == "*" {
			return true
		}
	}
	
	return false
}
