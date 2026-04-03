package utils

import "github.com/gin-gonic/gin"

func Success(message string, data interface{}) gin.H {
	return gin.H{
		"success": true,
		"message": message,
		"data":    data,
	}
}

func Error(message string, errs interface{}) gin.H {
	return gin.H{
		"success": false,
		"message": message,
		"errors":  errs,
	}
}
