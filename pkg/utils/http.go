package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ahmetilboga2004/go-blog/internal/middlewares"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleError(w http.ResponseWriter, status int, err error) {
	errorResponse := ErrorResponse{
		Code:    status,
		Message: err.Error(),
	}
	ResponseJSON(w, status, errorResponse)
}

func ResponseJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GetUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	rawUserID := r.Context().Value(middlewares.UserIDKey)

	// UUID tipinde ise direkt dön
	if userId, ok := rawUserID.(uuid.UUID); ok {
		return userId, nil
	}

	// String ise parse et
	if userIdStr, ok := rawUserID.(string); ok {
		return uuid.Parse(userIdStr)
	}

	return uuid.UUID{}, fmt.Errorf("invalid user ID format in context")
}
