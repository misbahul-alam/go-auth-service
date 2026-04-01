package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/database"
	"github.com/misbahul-alam/go-auth-service/internal/handler"
	"github.com/misbahul-alam/go-auth-service/internal/repository"
	"github.com/misbahul-alam/go-auth-service/internal/service"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	userRepo := repository.NewUserRepository(database.DB)
	authService := service.NewAuthService(userRepo)
	userHandler := handler.NewAuthHandler(authService)

	RegisterAuthRoutes(api, userHandler)
}
