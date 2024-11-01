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
	jwtService := services.NewJWTService(
		config.JWT.SecretKey,
		config.JWT.TokenExpiration,
		config.JWT.ResetTokenExpiration,
		config.JWT.VerificationTokenExpiration,
	)
	redisService := services.NewRedisService("localhost:6379", "", 0)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo, jwtService, redisService)
	userHandler := handlers.NewUserHandler(userService)

	postRepo := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	commentRepo := repository.NewCommentRepository(db)
	commentService := services.NewcommentService(commentRepo)
	commentHandler := handlers.NewCommentHandler(commentService)

	authMiddleware := middlewares.NewAuthMiddleware(jwtService, redisService)
	mux := http.NewServeMux()

	authMux := authMiddleware.Auth(mux)

	mux.HandleFunc("GET /users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("POST /users/register", authMiddleware.GuestOnly(userHandler.Register))
	mux.HandleFunc("POST /users/login", authMiddleware.GuestOnly(userHandler.Login))
	mux.HandleFunc("GET /users/logout", authMiddleware.RequireLogin(userHandler.Logout))

	mux.HandleFunc("GET /posts", postHandler.GetAllPosts)
	mux.HandleFunc("GET /posts/{id}", postHandler.GetPostByID)
	mux.HandleFunc("POST /posts", authMiddleware.RequireLogin(postHandler.Create))
	mux.HandleFunc("PUT /posts/{id}", authMiddleware.RequireLogin(postHandler.UpdatePost))
	mux.HandleFunc("DELETE /posts/{id}", authMiddleware.RequireLogin(postHandler.DeletePost))

	mux.HandleFunc("GET /comments", commentHandler.GetAllComments)
	mux.HandleFunc("GET /comments/{id}", commentHandler.GetCommentByID)
	mux.HandleFunc("POST /comments", commentHandler.Create)
	mux.HandleFunc("PUT /comments/{id}", commentHandler.UpdateComment)
	mux.HandleFunc("DELETE /comments/{id}", commentHandler.DeleteComment)

	server := &http.Server{
		Addr:    ":4000",
		Handler: authMux,
	}
	utils.Log(utils.INFO, "Sunucu başlatılıyor...")
	err := server.ListenAndServe()
	if err != nil {
		utils.Log(utils.ERROR, "Sunucu başlatılırken bir hata oluştu: %v", err)
	}

}
