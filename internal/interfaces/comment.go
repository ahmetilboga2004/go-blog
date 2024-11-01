package interfaces

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type CommentRepository interface {
	Create(comment *models.Comment) (*models.Comment, error)
	GetAll() ([]*models.Comment, error)
	GetByID(id uuid.UUID) (*models.Comment, error)
	Update(id uuid.UUID, comment *models.Comment) (*models.Comment, error)
	Delete(id uuid.UUID) error
}

type CommentService interface {
	CreateComment(userId uuid.UUID, comment *models.Comment) (*models.Comment, error)
	GetCommentByID(id uuid.UUID) (*models.Comment, error)
	GetAllComments() ([]*models.Comment, error)
	UpdateComment(userId, commentId uuid.UUID, comment *models.Comment) (*models.Comment, error)
	DeleteComment(userId, commentId uuid.UUID) error
}
