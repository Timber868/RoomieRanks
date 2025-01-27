package utils

import (
	"net/http"

	"github.com/rs/cors"
)

// CORS middleware function to allow CORS requests from specific origins
func CORS(next http.Handler) http.Handler {
	// Define the allowed origins and other CORS settings
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Apply the CORS middleware to the handler
	return corsHandler.Handler(next)
}
