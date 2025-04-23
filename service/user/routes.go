package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jakefish18/GolangHttpServer/service/auth"
	"github.com/jakefish18/GolangHttpServer/types"
	"github.com/jakefish18/GolangHttpServer/utils"
	"net/http"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handlerRegiser).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handlerRegiser(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterUserPaylod

	if err := utils.ParseJson(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetUserByEmail(payload.Email)

	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
