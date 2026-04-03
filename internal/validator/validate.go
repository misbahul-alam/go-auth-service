package validator

import "github.com/gin-gonic/gin"

func ValidateRequest(c *gin.Context, dto interface{}) (bool, map[string]string) {
	if err := c.ShouldBindJSON(dto); err != nil {
		return false, FormatErrors(err)
	}
	return true, nil
}
