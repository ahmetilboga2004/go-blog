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

func (h *commentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var commentReq dto.CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	comment := commentReq.ToModel()
	createdComment, err := h.commentService.CreateComment(userId, comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commentRes := dto.CommentResponseFromModel(createdComment)
	utils.ResJSON(w, http.StatusCreated, commentRes)
}

func (h *commentHandler) GetCommentByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comment, err := h.commentService.GetCommentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ResJSON(w, http.StatusOK, comment)
}

func (h *commentHandler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.commentService.GetAllComments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	commentsRes := dto.CommentListResponse(comments)
	utils.ResJSON(w, http.StatusOK, commentsRes)
}

func (h *commentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	commentIdStr := r.PathValue("id")
	commentId, err := uuid.Parse(commentIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var commentReq dto.CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validator.Struct(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	comment := commentReq.ToModel()

	updatedComment, err := h.commentService.UpdateComment(userId, commentId, comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resUpdatedComment := dto.CommentResponseFromModel(updatedComment)
	utils.ResJSON(w, http.StatusOK, resUpdatedComment)
}

func (h *commentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := h.commentService.DeleteComment(userId, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResJSON(w, http.StatusNoContent, "")
}
