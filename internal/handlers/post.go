package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmetilboga2004/go-blog/internal/dto"
	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/internal/middlewares"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type postHandler struct {
	postService interfaces.PostService
	validator   *validator.Validate
}

func NewPostHandler(postService interfaces.PostService) *postHandler {
	return &postHandler{
		postService: postService,
		validator:   validator.New(),
	}
}

func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	var postReq dto.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&postReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&postReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post := postReq.ToModel()
	rawUserID := r.Context().Value(middlewares.UserIDKey)
	userId, ok := rawUserID.(uuid.UUID)
	if !ok {
		if userIdStr, ok := rawUserID.(string); ok {
			parsedUUID, err := uuid.Parse(userIdStr)
			if err != nil {
				http.Error(w, "Invalid userId format in context", http.StatusInternalServerError)
				return
			}
			userId = parsedUUID
		} else {
			http.Error(w, "userId not found in context", http.StatusUnauthorized)
			return
		}
	}
	createdPost, err := h.postService.CreatePost(userId, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postRes := dto.PostResponseWithUserFromModel(createdPost)
	utils.ResJSON(w, http.StatusOK, postRes)
}
