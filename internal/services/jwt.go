package services

import (
	"errors"
	"time"

	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey                   string
	tokenExpiration             time.Duration
	resetTokenExpiration        time.Duration
	verificationTokenExpiration time.Duration
}

func NewJWTService(secretKey string, tokenExp, resetTokenExp, verificationTokenExp time.Duration) interfaces.JWTService {
	return &jwtService{
		secretKey:                   secretKey,
		tokenExpiration:             tokenExp,
		resetTokenExpiration:        resetTokenExp,
		verificationTokenExpiration: verificationTokenExp,
	}
}

func (s *jwtService) GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"exp":     time.Now().Add(s.tokenExpiration).Unix(),
	}
	return s.createTokenWithClaims(claims)
}

func (s *jwtService) GenerateEmailVerificationToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(s.verificationTokenExpiration).Unix(),
	}
	return s.createTokenWithClaims(claims)
}

func (s *jwtService) GeneratePasswordResetToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(s.resetTokenExpiration).Unix(),
	}
	return s.createTokenWithClaims(claims)
}

func (s *jwtService) ValidateToken(token string) (string, error) {
	claims, err := s.parseTokenClaims(token)
	if err != nil {
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid token payload")
	}
	return userID, nil
}

func (s *jwtService) ValidateEmailVerificationToken(token string) (string, error) {
	claims, err := s.parseTokenClaims(token)
	if err != nil {
		return "", err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid token payload")
	}
	return email, nil
}

func (s *jwtService) ValidatePasswordResetToken(token string) (string, error) {
	claims, err := s.parseTokenClaims(token)
	if err != nil {
		return "", err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid token payload")
	}
	return email, nil
}

func (s *jwtService) createTokenWithClaims(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) parseTokenClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
