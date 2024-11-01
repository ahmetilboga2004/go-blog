package services

import (
	"errors"
	"time"

	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
	"github.com/google/uuid"
)

type userService struct {
	userRepo     interfaces.UserRepository
	jwtService   interfaces.JWTService
	redisService interfaces.RedisService
	// mailService interfaces.MailService
}

// mailService interfaces.MailService
func NewUserService(userRepo interfaces.UserRepository, jwtService interfaces.JWTService, redisService interfaces.RedisService) interfaces.UserService {
	return &userService{
		userRepo:     userRepo,
		jwtService:   jwtService,
		redisService: redisService,
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
	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) LogoutUser(token string) error {
	claims, err := s.jwtService.ParseTokenClaims(token)
	if err != nil {
		return errors.New("invalid token")
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return errors.New("token expiration failed")
	}

	expiration := time.Until(time.Unix(int64(exp), 0))
	if expiration <= 0 {
		return errors.New("token expiration failed")
	}

	return s.redisService.BlacklistToken(token, expiration)
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
