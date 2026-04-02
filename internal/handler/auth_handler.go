package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/dto"
	"github.com/misbahul-alam/go-auth-service/internal/service"
	"github.com/misbahul-alam/go-auth-service/internal/utils"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.service.Login(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	AccessToken := utils.GenerateAccessToken(user.ID, string(user.Role))
	RefreshToken := utils.GenerateRefreshToken(user.ID, string(user.Role))

	c.JSON(200, gin.H{
		"access_token":  AccessToken,
		"refresh_token": RefreshToken,
		"type":          "Bearer",
	})
}
