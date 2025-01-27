package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Timber868/roomieranks/service/chore"
	"github.com/Timber868/roomieranks/service/chore_instance"
	"github.com/Timber868/roomieranks/service/household"
	"github.com/Timber868/roomieranks/service/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors" // Import the CORS package
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

	// Household routes
	householdStore := household.NewStore(s.db)
	householdHandler := household.NewHandler(householdStore)
	householdHandler.RegisterRoute(subrouter)

	// User routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore, householdStore)
	userHandler.RegisterRoute(subrouter)

	// Chore Routes
	choreStore := chore.NewStore(s.db)
	choreHandler := chore.NewHandler(choreStore)
	choreHandler.RegisterRoute(subrouter)

	// Chore-instance Routes
	choreInstanceStore := chore_instance.NewStore(s.db, choreStore) // <-- from chore_instance
	choreInstanceHandler := chore_instance.NewHandler(choreInstanceStore)
	choreInstanceHandler.RegisterRoute(subrouter)

	// Apply CORS middleware
	// Allow all origins by default, or you can specify a list of allowed origins
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},         // You can replace this with your frontend's URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},  // Allowed HTTP methods
		AllowedHeaders: []string{"Content-Type", "Authorization"}, // Allowed headers
	})

	// Wrap the router with CORS middleware
	handler := corsHandler.Handler(router)

	log.Println("Listening on", s.addr)

	// Start the server with the CORS-wrapped router
	return http.ListenAndServe(s.addr, handler)
}
