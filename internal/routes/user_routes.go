package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/handler"
	"github.com/misbahul-alam/go-auth-service/internal/middleware"
)

func RegisterUserRotes(r *gin.RouterGroup, h *handler.UserHandler) {
	users := r.Group("/users")

	users.Use(middleware.AuthMiddleware())

	users.GET("/me", h.GetMe)
}
