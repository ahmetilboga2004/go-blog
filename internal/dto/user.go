package dto

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/google/uuid"
)

type UserRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Username  string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

// containsany=!@#$%^&*"

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

func (r *UserRequest) ToModel() *models.User {
	return &models.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Username:  r.Username,
		Email:     r.Email,
		Password:  r.Password,
	}
}

func UserResponseFromModel(user *models.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
	}
}

func UserListResponse(users []*models.User) []*UserResponse {
	responses := make([]*UserResponse, len(users))
	for i, user := range users {
		responses[i] = UserResponseFromModel(user)
	}
	return responses
}
