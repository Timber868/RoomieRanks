package chore_instance

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
	store types.ChoreInstanceStore
}

func NewHandler(store types.ChoreInstanceStore) *Handler {
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

	router.HandleFunc("/chore-instance", h.handleCreateChoreInstance).Methods("POST")
	router.HandleFunc("/chore-instance/{id}", h.handleGetChoreInstanceByID).Methods("GET")
	router.HandleFunc("/chore-instance/assign/{id}", h.handleAssignChoreInstance).Methods("PUT")
	router.HandleFunc("/chore-instance/complete/{id}", h.handleCompleteChoreInstance).Methods("PUT")
}

func (h *Handler) handleCreateChoreInstance(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	//check if the chore exists if not create a new chore

	//Type we will use to decode our payload
	var payload types.RegisterChoreInstancePayload

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
	if err := h.store.CreateChoreInstance(types.ChoreInstance{
		Username: payload.Username,
		ChoreID:  payload.ChoreID,
		//DueDate is today's date plus seven days
		DueDate:   utils.GetNextWeek(),
		Completed: false,
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) handleGetChoreInstanceByID(w http.ResponseWriter, r *http.Request) {
	//get the id from the url
	vars := mux.Vars(r)
	id := vars["id"]

	//convert the id to an int
	idInt, err := utils.ConvertToInt(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Get the chore
	choreInstance, err := h.store.GetChoreInstanceByID(idInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, choreInstance)
}

func (h *Handler) handleAssignChoreInstance(w http.ResponseWriter, r *http.Request) {
	//get the id from the url
	vars := mux.Vars(r)
	id := vars["id"]

	//get the username from the request
	username, err := utils.GetUsernameFromRequest(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//convert the id to an int
	idInt, err := utils.ConvertToInt(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Assign the chore
	if err := h.store.AssignChoreInstance(idInt, username); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) handleCompleteChoreInstance(w http.ResponseWriter, r *http.Request) {
	//get the id from the url
	vars := mux.Vars(r)
	id := vars["id"]

	//convert the id to an int
	idInt, err := utils.ConvertToInt(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//Complete the chore
	if err := h.store.CompleteChore(idInt); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, "Chore Completed")
}
