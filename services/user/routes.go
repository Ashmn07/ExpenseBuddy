package user

import (
	"fmt"
	"net/http"

	"github.com/Ashmn07/ExpenseBuddy/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Login page")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Signup page")
}
