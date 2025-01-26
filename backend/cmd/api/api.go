package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Timber868/roomieranks/service/chore"
	chore_instance "github.com/Timber868/roomieranks/service/chore-instance"
	"github.com/Timber868/roomieranks/service/household"
	"github.com/Timber868/roomieranks/service/user"
	"github.com/gorilla/mux"
)

// Stuct that defines the API server it has both its address and also a pointer to the database
type APIServer struct {
	addr string
	db   *sql.DB
}

// Function that creates a new API server struct based on inputs
func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	//Api v1 subrouter
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// User routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoute(subrouter)

	// Household routes
	householdStore := household.NewStore(s.db)
	householdHandler := household.NewHandler(householdStore)
	householdHandler.RegisterRoute(subrouter)

	// Chore Routes
	choreStore := chore.NewStore(s.db)
	choreHandler := chore.NewHandler(choreStore)
	choreHandler.RegisterRoute(subrouter)

	// Chore Instance Routes
	choreInstanceStore := chore_instance.NewStore(s.db, choreStore)
	choreInstanceHandler := chore_instance.NewHandler(choreInstanceStore)
	choreInstanceHandler.RegisterRoute(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
