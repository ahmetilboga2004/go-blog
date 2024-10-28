package services

import (
	"errors"

	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/ahmetilboga2004/go-blog/internal/repository"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) RegisterUser(user *models.User) (*models.User, error) {
	existingUser, err := s.UserRepo.FindByUsernameOrEmail(user.Username, user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username or email already taken")
	}

	user, err = s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) LoginUser(usernameOrEmail, password string) (string, error) {
	user, err := s.UserRepo.FindByUsernameOrEmail(usernameOrEmail, usernameOrEmail)
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
