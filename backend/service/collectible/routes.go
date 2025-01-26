package collectible

import (
	"fmt"
	"net/http"

	"github.com/Timber868/roomieranks/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.CollectibleStore
}

func NewHandler(store types.CollectibleStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/collectible", h.handleCreateCollectible).Methods("POST")
}

func (h *Handler) handleCreateCollectible(w http.ResponseWriter, r *http.Request) {
	//Get the id from the url
	err := h.store.CreateCollectible("marrec")
	fmt.Println(err)
}
