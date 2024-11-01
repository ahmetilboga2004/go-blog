package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmetilboga2004/go-blog/internal/dto"
	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
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

	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	post := postReq.ToModel()
	createdPost, err := h.postService.CreatePost(userId, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postRes := dto.PostResponseWithUserFromModel(createdPost)
	utils.ResJSON(w, http.StatusOK, postRes)
}

func (h *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post, err := h.postService.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ResJSON(w, http.StatusOK, post)
}

func (h *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResJSON(w, http.StatusOK, posts)
}

func (h *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	postIdStr := r.PathValue("id")
	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var postReq dto.PostRequest
	if err := json.NewDecoder(r.Body).Decode(&postReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&postReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	post := postReq.ToModel()

	updatedPost, err := h.postService.UpdatePost(userId, postId, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resUpdatedPost := dto.PostResponseWithUserFromModel(updatedPost)

	utils.ResJSON(w, http.StatusOK, resUpdatedPost)
}

func (h *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.postService.DeletePost(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResJSON(w, http.StatusNoContent, "")
}
