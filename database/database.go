package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

// Connect to the database.
// FYI: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Set Data Source Name.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	// Set dialector.
	dialector := postgres.Open(dsn)
	// Open initialize db session based on dialector.
	Database, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		// The panic built-in function stops normal execution of the current goroutine.
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
