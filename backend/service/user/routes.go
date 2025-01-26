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
	router.HandleFunc("/user/{username}", h.handleGetUser).Methods("GET")
	router.HandleFunc("/user/{username}/level", h.handleLevelUp).Methods("POST")
	// router.HandleFunc("/user/{username}/title", h.handleChangeTitle).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	//Type we will use to decode our payload
	var payload types.LoginUserPayload

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

	u, err := h.store.GetUserByUsername(payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error, user not found, invalid username"))
		return
	}

	if !auth.ComparePasswords(u.Password, payload.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("error, invalid password"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Login successful")
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
	u, _ := h.store.GetUserByUsername(payload.Username)
	if u != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s already exists", payload.Username))
		return
	}

	//Check if the email exists
	u, _ = h.store.GetUserByEmail(payload.Email)
	if u != nil {
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
		Username:    payload.Username,
		Name:        payload.Name,
		Password:    hashedPassword,
		Email:       payload.Email,
		HouseholdID: 0, // Users dont start off with households
		Title:       "Newbie",
		Level:       1, //You start off at level 1
	})

	//Do error handling to validate data
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//DIsplay some message to send back to the user
	utils.WriteJSON(w, http.StatusCreated, "User created successfully")
}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	//Get the username from the url
	vars := mux.Vars(r)
	username, ok := vars["username"]

	// Validate that the username is there
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing username"))
		return
	}

	//Get the user from the store
	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleLevelUp(w http.ResponseWriter, r *http.Request) {
	//Get the username from the url
	vars := mux.Vars(r)
	username, ok := vars["username"]

	// Validate that the username is there
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing username"))
		return
	}

	//Level up the user
	err := h.store.LevelUp(username)
	if err == fmt.Errorf("user not found") {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, "User leveled up")
}
