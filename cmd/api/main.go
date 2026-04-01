package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misbahul-alam/go-auth-service/internal/config"
	"github.com/misbahul-alam/go-auth-service/internal/database"
	"github.com/misbahul-alam/go-auth-service/internal/model"
	"github.com/misbahul-alam/go-auth-service/internal/routes"
)

func main() {

	config.LoadEnv()
	database.Connect()
	_ = database.DB.AutoMigrate(&model.User{})

	r := gin.Default()

	routes.RegisterRoutes(r)

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}

}
