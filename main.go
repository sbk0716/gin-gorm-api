package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/middleware"
	"diary_api/model"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/*
main function:

1. Executes loadEnv function.

2. Executes loadDatabase function.

3. Executes serveApplication function.
*/
func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

/*
loadEnv function:

1. Reads env file and loads them into ENV for this process.
*/
func loadEnv() {
	const targetEnvName = "GO_ENV"
	if "" == os.Getenv(targetEnvName) {
		_ = os.Setenv(targetEnvName, "local")
	}
	filePath := fmt.Sprintf(".env.%s", os.Getenv(targetEnvName))
	fmt.Printf("filePath: %#v\n", filePath)
	// Reads env file and loads them into ENV for this process.
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv(targetEnvName))
	}
}

/*
loadDatabase function:

1. Opens the connection using the GORM PostgreSQL driver.

2. Runs auto migration for given models.
*/
func loadDatabase() {
	// Opens the connection using the GORM PostgreSQL driver.
	database.Connect()

	// Runs auto migration for given models.
	// FYI: https://gorm.io/docs/migration.html
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

/*
serveApplication function:

1. Returns an Engine instance with the Logger and Recovery middleware already attached.

2. Creates a new router group(publicRoutes).

3. Creates a new router group(protectedRoutes) with additional custom middleware(JWTAuthMiddleware).

4. Attaches the router to a http.Server and starts listening and serving HTTP requests.
*/
func serveApplication() {
	// Returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()

	// Creates a new router group(publicRoutes).
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	// Creates a new router group(protectedRoutes) with additional custom middleware(JWTAuthMiddleware).
	protectedRoutes := router.Group("/api")
	// Adds middleware to the group.
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)
	protectedRoutes.GET("/entry/:id", controller.GetEntry)

	// Attaches the router to a http.Server and starts listening and serving HTTP requests.
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
