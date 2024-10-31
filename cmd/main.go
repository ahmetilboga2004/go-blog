package main

import (
	"net/http"

	"github.com/ahmetilboga2004/go-blog/config"
	"github.com/ahmetilboga2004/go-blog/config/database"
	"github.com/ahmetilboga2004/go-blog/internal/handlers"
	"github.com/ahmetilboga2004/go-blog/internal/middlewares"
	"github.com/ahmetilboga2004/go-blog/internal/repository"
	"github.com/ahmetilboga2004/go-blog/internal/services"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
)

func main() {
	db := database.InitDB()

	config.LoadConfig()
	userRepo := repository.NewUserRepository(db)
	jwtService := services.NewJWTService(
		config.JWT.SecretKey,
		config.JWT.TokenExpiration,
		config.JWT.ResetTokenExpiration,
		config.JWT.VerificationTokenExpiration,
	)
	redisService := services.NewRedisService("localhost:6379", "", 0)
	userService := services.NewUserService(userRepo, jwtService, redisService)
	userHandler := handlers.NewUserHandler(userService)

	authMiddleware := middlewares.NewAuthMiddleware(jwtService)
	mux := http.NewServeMux()

	authMux := authMiddleware.Auth(mux)

	mux.HandleFunc("POST /users/register", authMiddleware.GuestOnly(userHandler.Register))
	mux.HandleFunc("POST /users/login", userHandler.Login)

	server := &http.Server{
		Addr:    ":4000",
		Handler: authMux,
	}
	err := server.ListenAndServe()
	if err != nil {
		utils.Log(utils.ERROR, "Sunucu başlatılırken bir hata oluştu %v", err)
	} else {
		utils.Log(utils.INFO, "Sunucu başlatıldı")
	}
}
