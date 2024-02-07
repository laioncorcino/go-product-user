package handler

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/laioncorcino/go-product-user/internal/dto"
	"github.com/laioncorcino/go-product-user/internal/entity"
	"github.com/laioncorcino/go-product-user/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB database.UserQuery
}

func NewUserHandler(userDB database.UserQuery) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request dto.UserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(request.Name, request.Email, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(request.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if validPass := user.ValidatePass(request.Password); validPass == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	expires := r.Context().Value("expires").(int)

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub":  user.UserID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Second * time.Duration(expires)).Unix(),
	})

	tokenResponse := dto.TokenResponse{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokenResponse)
}
