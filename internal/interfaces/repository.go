package interfaces

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetAll() ([]*models.User, error)
	GetByID(id uuid.UUID) (*models.User, error)
	FindByUsernameOrEmail(username, email string) (*models.User, error)
	Update(id uuid.UUID, user *models.User) (*models.User, error)
	Delete(id uuid.UUID) error
	// UpdatePassword(id uuid.UUID, hashedPassword, salt string) error
}
