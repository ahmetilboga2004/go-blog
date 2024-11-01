package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ahmetilboga2004/go-blog/internal/dto"
	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	utils.ResJSON(w, http.StatusOK, userRes)
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

func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Token yok", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if err := h.userService.LogoutUser(token); err != nil {
		http.Error(w, "Logout işlemi başarısız", http.StatusInternalServerError)
		return
	}

	utils.ResJSON(w, http.StatusOK, "Çıkış Başarılı")
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	usersRes := dto.UserListResponse(users)
	utils.ResJSON(w, http.StatusOK, usersRes)
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userRes := dto.UserResponseFromModel(user)
	utils.ResJSON(w, http.StatusOK, userRes)
}
