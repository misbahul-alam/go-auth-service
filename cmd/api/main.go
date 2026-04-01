package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/config"
	"github.com/misbahul-alam/go-auth-service/internal/database"
	"github.com/misbahul-alam/go-auth-service/internal/model"
)

func main() {

	config.LoadEnv()
	database.Connect()
	_ = database.DB.AutoMigrate(&model.User{})

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}

}
