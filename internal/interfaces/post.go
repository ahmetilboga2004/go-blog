package interfaces

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type PostRepository interface {
	Create(post *models.Post) (*models.Post, error)
	GetAll() ([]*models.Post, error)
	GetByID(id uuid.UUID) (*models.Post, error)
	Update(id uuid.UUID, post *models.Post) (*models.Post, error)
	Delete(id uuid.UUID) error
}

type PostService interface {
	CreatePost(userId uuid.UUID, post *models.Post) (*models.Post, error)
	GetPostByID(id uuid.UUID) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
	UpdatePost(userId, postId uuid.UUID, post *models.Post) (*models.Post, error)
	DeletePost(userId, postId uuid.UUID) error
}
