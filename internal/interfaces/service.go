package interfaces

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
)

type UserService interface {
	RegisterUser(user *models.User) (*models.User, error)
	LoginUser(usernameOrEmail, password string) (string, error)
	LogoutUser(token string) error
	// GetUserByID(id uuid.UUID) (*models.User, error)
	// UpdateUser(id uuid.UUID, user *models.User) (*models.User, error)
	// DeleteUser(id uuid.UUID) error
	// ChangePassword(id uuid.UUID, oldPassword, newPassword string) error
	// ResetPassword(email string) error
	// VerifyEmail(token string) error
	// GetAllUsers() ([]*models.User, error)
}
