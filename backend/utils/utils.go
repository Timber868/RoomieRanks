package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Does caching so you only want to instatniate it once
var Validate = validator.New() //Users all the types validate section

// Helper function that we are going to use over and over again
func ParseJson(r *http.Request, payload any) error {
	//Make sure body is not empy to avoid nil pointers
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader((status))

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
