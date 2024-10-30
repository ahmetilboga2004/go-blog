package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmetilboga2004/go-blog/internal/dto"
	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService interfaces.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService interfaces.UserService) *userHandler {
	return &userHandler{
		userService: userService,
		validator:   validator.New(),
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validator.Struct(&userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := userReq.ToModel()
	createdUser, err := h.userService.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRes := dto.UserResponseFromModel(createdUser)
	json.NewEncoder(w).Encode(userRes)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		UsernameOrEmail string `json:"username_or_email"`
		Password        string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.userService.LoginUser(creds.UsernameOrEmail, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responseData := map[string]string{"token": token}
	utils.ResJSON(w, http.StatusOK, responseData)
}
