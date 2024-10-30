package services

import (
	"errors"

	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
)

type userService struct {
	userRepo interfaces.UserRepository
	// mailService interfaces.MailService
	// jwtService interfaces.JWTService
}

// mailService interfaces.MailService
// jwtService interfaces.JWTService
func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) RegisterUser(user *models.User) (*models.User, error) {
	existingUser, err := s.userRepo.FindByUsernameOrEmail(user.Username, user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username or email already taken")
	}

	user, err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) LoginUser(usernameOrEmail, password string) (string, error) {
	user, err := s.userRepo.FindByUsernameOrEmail(usernameOrEmail, usernameOrEmail)
	if err != nil || user == nil {
		return "", errors.New("invalid username or email")
	}
	hashedPassword := utils.HashPassword(password, user.Salt)
	if hashedPassword != user.Password {
		return "", errors.New("invalid password")
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
