package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/misbahul-alam/go-auth-service/docs"
	"github.com/misbahul-alam/go-auth-service/internal/config"
	"github.com/misbahul-alam/go-auth-service/internal/database"
	"github.com/misbahul-alam/go-auth-service/internal/model"
	"github.com/misbahul-alam/go-auth-service/internal/routes"
	"github.com/misbahul-alam/go-auth-service/internal/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Auth Service
// @version 1.0
// @description A simple authentication backend built with Golang, Gin, and GORM. This project demonstrates core backend concepts like JWT authentication, password hashing, and clean structure.
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type 'Bearer ' followed by your JWT token.

// @host localhost:8080
// @BasePath /api/v1
func main() {

	config.LoadEnv()
	database.Connect()
	validator.Init()
	_ = database.DB.AutoMigrate(&model.User{})

	r := gin.Default()

	routes.RegisterRoutes(r)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")

	if err != nil {
		panic(err)
	}

}
