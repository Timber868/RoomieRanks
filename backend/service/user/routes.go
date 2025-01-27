package user

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Timber868/roomieranks/service/auth"
	"github.com/Timber868/roomieranks/service/household"
	"github.com/Timber868/roomieranks/types"
	"github.com/Timber868/roomieranks/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Each service has a handler
type Handler struct {
	userStore      types.UserStore
	householdStore *household.Store
}

func NewHandler(userStore types.UserStore, householdStore *household.Store) *Handler {
	return &Handler{
		userStore:      userStore,
		householdStore: householdStore,
	}
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	// CORS middleware: Only allow requests from localhost:5173 (adjust the port if needed)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Allow only this origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Apply the CORS handler to the router
	router.Use(corsHandler.Handler)

	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/user/{username}", h.handleGetUser).Methods("GET")
	router.HandleFunc("/user/{username}/household/{householdID}", h.handleChangeHousehold).Methods("PUT")
	router.HandleFunc("/user/{username}/xp", h.handleAddXP).Methods("PUT")
	router.HandleFunc("/user/{username}/title", h.handleChangeTitle).Methods("PUT")
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

	u, err := h.userStore.GetUserByUsername(payload.Username)
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

	// Check if the user already exists
	_, err := h.userStore.GetUserByUsername(payload.Username)
	// If err == nil, user was actually found -> can’t register
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s already exists", payload.Username))
		return
	}

	// Same logic for the email check
	_, err = h.userStore.GetUserByEmail(payload.Email)
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
	err = h.userStore.CreateUser(types.User{
		Username:    payload.Username,
		Name:        payload.Name,
		Password:    hashedPassword,
		Email:       payload.Email,
		HouseholdID: 0, // Users dont start off with households
		Title:       "Newbie",
		Level:       1, //You start off at level 1
		XP:          0, //You start off with 0 xp
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
	user, err := h.userStore.GetUserByUsername(username)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleAddXP(w http.ResponseWriter, r *http.Request) {
	//Get the username from the url
	vars := mux.Vars(r)
	username, ok := vars["username"]

	// Validate that the username is there
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing username"))
		return
	}

	var payload types.AddXPPayload

	//Make sure it is a valid json
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

	//Level up the user
	err := h.userStore.AddXP(username, payload.XP)
	if err != nil && strings.Contains(err.Error(), "user not found") {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "XP added successfully")
}

func (h *Handler) handleChangeTitle(w http.ResponseWriter, r *http.Request) {
	//Get the username from the url
	vars := mux.Vars(r)
	username, ok := vars["username"]

	// Validate that the username is there
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing username"))
		return
	}

	//Type we will use to decode our payload
	var payload types.ChangeTitlePayload

	//Make sure it is a valid json
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

	//Change the title
	err := h.userStore.ChangeTitle(username, payload.Title)
	if err == fmt.Errorf("user not found") {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, "Title changed")
}

/*
Method to change the household of a user
If you use the id 0 it unsets it
If not the household needs to exist
*/
func (h *Handler) handleChangeHousehold(w http.ResponseWriter, r *http.Request) {
	//Get the username from the url
	vars := mux.Vars(r)
	username, ok := vars["username"]

	// Validate that the username is there
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing username"))
		return
	}

	//Get the householdID from the url
	household, ok := vars["householdID"]

	// Validate that the username is there
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing householdID"))
		return
	}
	householdID, err := strconv.Atoi(household)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Check that the household we want to change to exists
	u, _ := h.householdStore.GetHouseholdByID(householdID)
	if u == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("household with id %d not found", householdID))
		return
	}

	//Change the household
	err = h.userStore.ChangeHousingID(username, householdID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Household changed")
}
