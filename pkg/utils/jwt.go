package utils

import (
	"time"

	"github.com/ahmetilboga2004/go-blog/config"
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *models.User) (string, error) {
	claims := &jwt.MapClaims{
		"userId": user.ID.String(),
		"exp":    time.Now().Add(config.TokenExpiryDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWTSecret)
}
