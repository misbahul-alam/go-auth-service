package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/handler"
)

func RegisterAuthRoutes(r *gin.RouterGroup, h *handler.AuthHandler) {
	auth := r.Group("/auth")

	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
}
