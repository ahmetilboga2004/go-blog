package interfaces

import (
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(user *models.User) (string, error)
	GenerateEmailVerificationToken(email string) (string, error)
	GeneratePasswordResetToken(email string) (string, error)
	ValidateToken(token string) (string, error)
	ValidateEmailVerificationToken(token string) (string, error)
	ValidatePasswordResetToken(token string) (string, error)
	CreateTokenWithClaims(claims jwt.MapClaims) (string, error)
	ParseTokenClaims(tokenStr string) (jwt.MapClaims, error)
}
