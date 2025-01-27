package collectible

import (
	"fmt"
	"net/http"

	"github.com/Timber868/roomieranks/types"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Handler struct {
	store types.CollectibleStore
}

func NewHandler(store types.CollectibleStore) *Handler {
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

	router.HandleFunc("/collectible", h.handleCreateCollectible).Methods("POST")
}

func (h *Handler) handleCreateCollectible(w http.ResponseWriter, r *http.Request) {
	//Get the id from the url
	err := h.store.CreateCollectible("marrec")
	fmt.Println(err)
}
