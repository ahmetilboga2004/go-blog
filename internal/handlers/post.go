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

// Create godoc
// @Tags posts
// @Accept json
// @Produce json
// @Summary Create a new post
// @Description Create a new post with the provided data
// @Param post body dto.PostReq true "Post bilgileri"
// @Success 201 {object} dto.PostResp
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /posts [post]
func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	var postReq dto.PostReq
	if err := json.NewDecoder(r.Body).Decode(&postReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(&postReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err)
		return
	}

	post := postReq.ToModel()
	createdPost, err := h.postService.CreatePost(userId, post)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	postRes := dto.FromPostDetail(createdPost)
	utils.ResponseJSON(w, http.StatusOK, postRes)
}

// GetPostByID godoc
// @Tags posts
// @Accept json
// @Produce json
// @Summary Get a post by ID
// @Description Retrieve a post by its unique ID
// @Param id path string true "Post ID"
// @Success 200 {object} dto.PostDetailResp
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /posts/{id} [get]
func (h *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	post, err := h.postService.GetPostByID(id)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, err)
		return
	}
	utils.ResponseJSON(w, http.StatusOK, post)
}

// GetAllPosts godoc
// @Tags posts
// @Accept json
// @Produce json
// @Summary Get all posts
// @Description Retrieve a list of all posts
// @Success 200 {array} dto.PostResp
// @Failure 400 {object} utils.ErrorResponse
// @Router /posts [get]
func (h *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.postService.GetAllPosts()
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	postsRes := dto.FromPostList(posts)
	utils.ResponseJSON(w, http.StatusOK, postsRes)
}

// UpdatePost godoc
// @Tags posts
// @Accept json
// @Produce json
// @Summary Update a post by ID
// @Description Update a post with the provided ID and data
// @Param id path string true "Post ID"
// @Param post body dto.PostReq true "Post bilgileri"
// @Success 200 {object} dto.PostResp
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /posts/{id} [put]
func (h *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	postIdStr := r.PathValue("id")
	postId, err := uuid.Parse(postIdStr)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	var postReq dto.PostReq
	if err := json.NewDecoder(r.Body).Decode(&postReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(&postReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err)
		return
	}

	post := postReq.ToModel()

	updatedPost, err := h.postService.UpdatePost(userId, postId, post)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	resUpdatedPost := dto.FromPostDetail(updatedPost)

	utils.ResponseJSON(w, http.StatusOK, resUpdatedPost)
}

// DeletePost godoc
// @Tags posts
// @Accept json
// @Produce json
// @Summary Delete a post by ID
// @Description Remove a post with the specified ID
// @Param id path string true "Post ID"
// @Success 204
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /posts/{id} [delete]
func (h *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err)
		return
	}
	if err := h.postService.DeletePost(userId, id); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	utils.ResponseJSON(w, http.StatusNoContent, "")
}
