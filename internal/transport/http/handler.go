package http

import (
	"encoding/json"
	"net/http"

	"github.com/FrancoBarrera99/auth-service/internal/auth"
	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
)

type Handler struct {
	serv auth.AuthService
}

func NewHandler(serv auth.AuthService) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var creds model.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, token, err := h.serv.Login(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	userRes := model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	response := map[string]interface{}{
		"token": token,
		"user":  userRes,
	}
	w.Header().Set("Conent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var reg model.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	token, err := h.serv.Register(reg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	response := map[string]string{"token": token}
	w.Header().Set("Conent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
