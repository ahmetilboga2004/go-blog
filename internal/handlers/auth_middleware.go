package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/ahmetilboga2004/go-blog/config"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDKey contextKey = "userId"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return config.JWTSecret, nil
		})

		if token != nil && token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				userID := claims["userId"].(string)
				ctx := context.WithValue(r.Context(), userIDKey, userID)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(userIDKey) == nil {
			http.Error(w, "You must be logged in to access this resource", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GuestOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(userIDKey) != nil {
			http.Error(w, "This resource is only accessible to guests", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
