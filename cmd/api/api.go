package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperionXCF/rest-api/cmd/services/user"
)

type APIServer struct {
	address string 
	db *sql.DB	
}

// create new instance of api server
// returns pointer to the api server
func NewApiServer(address string, db  *sql.DB) *APIServer {
	return &APIServer{
		address : address,
		db : db,
	}
}

// function on APIServer struct to run
// Run fn will create a router, register all routes
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// create the service functions / request handlers in another folder 
	// register them with router here
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}