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

type commentHandler struct {
	commentService interfaces.CommentService
	validator      *validator.Validate
}

func NewCommentHandler(commentService interfaces.CommentService) *commentHandler {
	return &commentHandler{
		commentService: commentService,
		validator:      validator.New(),
	}
}

// Create godoc
// @Tags comments
// @Accept json
// @Produce json
// @Summary Create a new comment
// @Description Create a new comment
// @Param comment body dto.CommentRequest true "Yorum bilgileri"
// @Success 201 {object} dto.CommentResponse
// @Success 401 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /comments [post]
func (h *commentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var commentReq dto.CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.validator.Struct(&commentReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err)
		return
	}
	comment := commentReq.ToModel()
	createdComment, err := h.commentService.CreateComment(userId, comment)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	commentRes := dto.CommentResponseFromModel(createdComment)
	utils.ResponseJSON(w, http.StatusCreated, commentRes)
}

// GetCommentByID godoc
// @Tags comments
// @Accept json
// @Produce json
// @summary Get a comment by id
// @Description Retrieve a comment by its unique ID
// @Param id path string true "Comment ID"
// @Success 200 {object} dto.CommentResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /comments/{id} [get]
func (h *commentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	comment, err := h.commentService.GetCommentByID(id)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, err)
		return
	}
	utils.ResponseJSON(w, http.StatusOK, comment)
}

// GetAllComments godoc
// @Tags comments
// @Accept json
// @Produce json
// @Summary Get all comments
// @Description Retrieve a list of all comments
// @Success 200 {array} dto.CommentResponse "Empty array if no comments"
// @Failure 400 {object} utils.ErrorResponse
// @Router /comments [get]
func (h *commentHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.commentService.GetAllComments()
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	commentsRes := dto.CommentListResponse(comments)
	utils.ResponseJSON(w, http.StatusOK, commentsRes)
}

// UpdateComment godoc
// @Tags comments
// @Accept json
// @Produce json
// @Summary Update a comment by ID
// @Description Update a comment with the provided ID
// @Param id path string true "Comment ID"
// @Param comment body dto.CommentRequest true "Yorum bilgileri"
// @Success 200 {object} dto.CommentResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /comments/{id} [put]
func (h *commentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	commentIdStr := r.PathValue("id")
	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	var commentReq dto.CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validator.Struct(&commentReq); err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}
	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.HandleError(w, http.StatusUnauthorized, err)
		return
	}
	comment := commentReq.ToModel()

	updatedComment, err := h.commentService.UpdateComment(userId, commentId, comment)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err)
		return
	}

	resUpdatedComment := dto.CommentResponseFromModel(updatedComment)
	utils.ResponseJSON(w, http.StatusOK, resUpdatedComment)
}

// DeleteComment godoc
// @Tags comments
// @Accept json
// @Produce json
// @Summary Delete a comment by ID
// @Description Remove a comment with the specified ID
// @Param id path string true "Comment ID"
// @Success 204
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /comments/{id} [delete]
func (h *commentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
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
	if err := h.commentService.DeleteComment(userId, id); err != nil {
		utils.HandleError(w, http.StatusNotFound, err)
		return
	}
	utils.ResponseJSON(w, http.StatusNoContent, "")
}
