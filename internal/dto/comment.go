package dto

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type CommentRequest struct {
	Content string    `json:"content" validate:"required"`
	PostID  uuid.UUID `json:"postId" validate:"required" format:"uuid"`
}

type CommentResponse struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"userId"`
	PostID  uuid.UUID `json:"postId"`
}

func (r *CommentRequest) ToModel() *models.Comment {
	return &models.Comment{
		Content: r.Content,
	}
}

func CommentResponseFromModel(comment *models.Comment) *CommentResponse {
	return &CommentResponse{
		ID:      comment.ID,
		Content: comment.Content,
		UserID:  comment.UserID,
		PostID:  comment.PostID,
	}
}

func CommentListResponse(comments []*models.Comment) []*CommentResponse {
	responses := make([]*CommentResponse, len(comments))
	for i, comment := range comments {
		responses[i] = CommentResponseFromModel(comment)
	}
	return responses
}
