package main

import (
	"net/http"

	"github.com/ahmetilboga2004/go-blog/config/database"
	"github.com/ahmetilboga2004/go-blog/internal/handlers"
	"github.com/ahmetilboga2004/go-blog/internal/repository"
	"github.com/ahmetilboga2004/go-blog/internal/services"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
)

func main() {
	db := database.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users/register", userHandler.Register)
	mux.HandleFunc("POST /users/login", userHandler.Login)

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		utils.Log(utils.ERROR, "Sunucu başlatılırken bir hata oluştu %v", err)
	} else {
		utils.Log(utils.INFO, "Sunucu başlatıldı")
	}
}
