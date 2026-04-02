package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetUint("user_id")
	user, err := h.service.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
