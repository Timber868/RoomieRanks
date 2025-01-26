package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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

func GetIDFromRequest(r *http.Request, key string) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars[key]
	if !ok {
		return 0, fmt.Errorf("missing id in request")
	}

	return strconv.Atoi(id)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ConvertToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func GetUsernameFromRequest(r *http.Request) (string, error) {
	//Get the username from the request
	username, ok := r.Context().Value("username").(string)
	if !ok {
		return "", fmt.Errorf("username not found in request")
	}

	return username, nil
}

func GetNextWeek() time.Time {
	return time.Now().AddDate(0, 0, 7)
}
