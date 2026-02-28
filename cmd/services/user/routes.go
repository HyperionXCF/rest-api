package user

import (
	"net/http"
	"github.com/gorilla/mux"
)

// struct for each router handler function
type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

// function on handler struct, register routes
func (h *Handler) RegisterRoutes(router *mux.Router) { 
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register" ,h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}


