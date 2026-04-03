package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/dto"
	"github.com/misbahul-alam/go-auth-service/internal/service"
	"github.com/misbahul-alam/go-auth-service/internal/utils"
	"github.com/misbahul-alam/go-auth-service/internal/validator"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if ok, errs := validator.ValidateRequest(c, &req); !ok {
		c.JSON(http.StatusBadRequest, utils.Error("validation failed", errs))
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

	if ok, errs := validator.ValidateRequest(c, &req); !ok {
		c.JSON(http.StatusBadRequest, utils.Error("validation failed", errs))
		return
	}

	AccessToken, RefreshToken, err := h.service.Login(req.Email, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  AccessToken,
		"refresh_token": RefreshToken,
		"type":          "Bearer",
	})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest

	if ok, errs := validator.ValidateRequest(c, &req); !ok {
		c.JSON(http.StatusBadRequest, utils.Error("validation failed", errs))
		return
	}
	AccessToken, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"access_token": AccessToken,
	})
}
