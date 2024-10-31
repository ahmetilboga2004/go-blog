package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
)

type contextKey string

const userIDKey contextKey = "userId"

type authMiddleware struct {
	jwtService   interfaces.JWTService
	redisService interfaces.RedisService
}

func NewAuthMiddleware(jwtService interfaces.JWTService, redisService interfaces.RedisService) *authMiddleware {
	return &authMiddleware{
		jwtService:   jwtService,
		redisService: redisService,
	}
}

func (m *authMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		isBlacklisted, err := m.redisService.IsBlacklistedToken(tokenString)
		if err != nil || isBlacklisted {
			next.ServeHTTP(w, r)
			return
		}

		userID, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(userIDKey) == nil {
			http.Error(w, "You must be logged in to access this resource", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *authMiddleware) GuestOnly(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(userIDKey) != nil {
			http.Error(w, "This resource is only accessible to guests", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
