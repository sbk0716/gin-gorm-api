package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	// Reads env file and loads them into ENV for this process.
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	// Runs auto migration for given model(User).
	database.Database.AutoMigrate(&model.User{})
	// Runs auto migration for given model(Entry).
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	// Returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()

	// Creates a new router group(publicRoutes).
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	// Creates a new router group(protectedRoutes).
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	// Attaches the router to a http.Server and starts listening and serving HTTP requests.
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
