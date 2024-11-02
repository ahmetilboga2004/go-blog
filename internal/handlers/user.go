package handlers

import (
	"encoding/json"
	"errors"
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

// @Summary Kullanıcı kaydı
// @Description Yeni bir kullanıcı oluşturur.
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserRequest true "Kullanıcı bilgileri"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/register [post]
func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	if err := h.validator.Struct(&userReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	user := userReq.ToModel()
	createdUser, err := h.userService.RegisterUser(user)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	userRes := dto.UserResponseFromModel(createdUser)
	utils.ResponseJSON(w, http.StatusOK, userRes)
}

// @Summary Kullanıcı girişi
// @Description Kullanıcının giriş yapmasını sağlar ve JWT token döner.
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Kullanıcı adı veya email ve şifre"
// @Success 200 {object} map[string]string
// @Failure 400 {object} utils.ErrorResponse
// @Router /users/login [post]
func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(creds); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.userService.LoginUser(creds.UsernameOrEmail, creds.Password)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	responseData := map[string]string{"token": token}
	utils.ResponseJSON(w, http.StatusOK, responseData)
}

// @Summary Kullanıcı çıkışı
// @Description Kullanıcının çıkış yapmasını sağlar.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "Çıkış Başarılı"
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /users/logout [get]
func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		utils.HandleError(w, http.StatusBadRequest, errors.New("unauthorized"))
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if err := h.userService.LogoutUser(token); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Çıkış Başarılı")
}

// @Summary Tüm kullanıcıları getir
// @Description Veritabanındaki tüm kullanıcıları listeler.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} dto.UserResponse "Empty array if no users"
// @Failure 500 {object} utils.ErrorResponse
// @Router /users [get]
func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	usersRes := dto.UserListResponse(users)
	utils.ResponseJSON(w, http.StatusOK, usersRes)
}

// @Summary Belirli bir kullanıcıyı getir
// @Description ID'sine göre kullanıcıyı getirir.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "Kullanıcı ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /users/{id} [get]
func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	user, err := h.userService.GetUserByID(id)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	userRes := dto.UserResponseFromModel(user)
	utils.ResponseJSON(w, http.StatusOK, userRes)
}
