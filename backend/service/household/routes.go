package household

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Timber868/roomieranks/types"
	"github.com/Timber868/roomieranks/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.HouseholdStore
}

// Retruns the handler our api will use
func NewHandler(store types.HouseholdStore) *Handler {
	return &Handler{store: store}
}

// Registers the routes for the api so it handles it when its called
func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/household", h.handleCreateHousehold).Methods("POST")
	router.HandleFunc("/household/{id}", h.handleGetHousehold).Methods("GET")
}

func (h *Handler) handleCreateHousehold(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterHouseholdPayload

	//IF we cant parse the json it might not be of the right type
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//Create the household
	err := h.store.CreateHousehold(types.Household{Name: payload.Name})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	message := fmt.Sprintf("Household with name %s created", payload.Name)

	// If everything went well we can return a message
	utils.WriteJSON(w, http.StatusCreated, message)
}

func (h *Handler) handleGetHousehold(w http.ResponseWriter, r *http.Request) {
	//Get the id from the url
	vars := mux.Vars(r)
	str, ok := vars["id"]

	fmt.Println("ID", str)

	//Validate that the id is there
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing id"))
		return
	}

	//Parse the id to an int
	householdID, err := strconv.Atoi(str)

	//If we cant parse the id we return an error as it should be a number
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Get the household from the store
	household, err := h.store.GetHouseholdByID(householdID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	//If we didnt find the household we return a 404
	if household == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("household with id %d not found", householdID))
		return
	}

	//If we found the household we return it
	utils.WriteJSON(w, http.StatusOK, household)
}
