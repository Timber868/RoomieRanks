package user

import (
	"fmt"
	"net/http"

	"github.com/Timber868/roomieranks/service/auth"
	"github.com/Timber868/roomieranks/types"
	"github.com/Timber868/roomieranks/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Each service has a handler
type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	//check if the user exists if not create a new user

	//Type we will use to decode our payload
	var payload types.RegisterUserPayload

	//Write it to json so that it can be sent through our api that there was an error
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors) //Necessary to get the erro message
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//Check in the database if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	//Hash the password for security reasons
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//If not then you can register the user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Password:  hashedPassword,
		Email:     payload.Email,
	})

	//Do error handling to validate data
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//DIsplay some message to send back to the user
	utils.WriteJSON(w, http.StatusCreated, "User created successfully")
}
