package chore

import (
	"fmt"
	"net/http"

	"github.com/Timber868/roomieranks/types"
	"github.com/Timber868/roomieranks/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Each service has a handler
type Handler struct {
	store types.ChoreStore
}

func NewHandler(store types.ChoreStore) *Handler {
	return &Handler{store: store}
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

	router.HandleFunc("/chore", h.handleCreateChore).Methods("POST")
	router.HandleFunc("/chore/assign", h.handleAssignChores).Methods("POST") // Method to assign all unassigned chores to users. Create chores_instance basically
	router.HandleFunc("/chore/{id}", h.handleGetChoreByID).Methods("GET")
	router.HandleFunc("/chore/household/{id}", h.handleGetChoreByHouseholdID).Methods("GET")
	router.HandleFunc("/chore/{id}", h.handleUpdateChore).Methods("PUT")
	router.HandleFunc("/chore/{id}", h.handleDeleteChore).Methods("DELETE")

}

func (h *Handler) handleCreateChore(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	//check if the chore exists if not create a new chore

	//Type we will use to decode our payload
	var payload types.RegisterChorePayload

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

	//Create the chore
	if err := h.store.CreateChore(types.Chore{
		Name:            payload.Name,
		Difficulty:      payload.Difficulty,
		TimeEstimate:    payload.TimeEstimate,
		CompletitonTime: payload.CompletitonTime,
		HouseholdID:     payload.HouseholdID,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetChoreByID(w http.ResponseWriter, r *http.Request) {
	//Get the id from the request
	id, err := utils.GetIDFromRequest(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Get the chore
	c, err := h.store.GetChoreByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, c)
}

func (h *Handler) handleGetChoreByHouseholdID(w http.ResponseWriter, r *http.Request) {
	//Get the id from the request
	id, err := utils.GetIDFromRequest(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Get the chore
	c, err := h.store.GetChoreByHouseholdID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, c)
}

func (h *Handler) handleUpdateChore(w http.ResponseWriter, r *http.Request) {
	//Get the id from the request
	id, err := utils.GetIDFromRequest(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Get the chore
	c, err := h.store.GetChoreByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//Type we will use to decode our payload
	var payload types.RegisterChorePayload

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

	//Update the chore
	if err := h.store.UpdateChore(types.Chore{
		ID:              id,
		Name:            payload.Name,
		Difficulty:      payload.Difficulty,
		TimeEstimate:    payload.TimeEstimate,
		CompletitonTime: payload.CompletitonTime,
		HouseholdID:     c.HouseholdID,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, nil)
}

func (h *Handler) handleDeleteChore(w http.ResponseWriter, r *http.Request) {
	//Get the id from the request
	id, err := utils.GetIDFromRequest(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Delete the chore
	if err := h.store.DeleteChore(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, nil)
}

func (h *Handler) handleAssignChores(w http.ResponseWriter, r *http.Request) {
	//Get the id from the request
	id, err := utils.GetIDFromRequest(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Get the chores
	_, err = h.store.GetChoreByHouseholdID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Add up the total difficulty of the chores
	// Find how many roomates are linked to that household
	// Divide the total difficulty by the number of roomates
	// Assign the chores to the roomates according to the difficulty
	// If there are any chores left over, assign them to the roomate with the least chores;p
}
