package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

/*
1. Retrieves the environment variables required to set up a database connection.
2. Opens the connection using the GORM PostgreSQL driver.
FYI: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
*/
func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Sets Data Source Name to dsn.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	// Sets gorm.Dialector to pgDialector.
	pgDialector := postgres.Open(dsn)
	// Sets options to pgOpts.
	pgOpts := &gorm.Config{}
	// gorm.Open function opens initialize db session based on dialector.
	Database, err = gorm.Open(pgDialector, pgOpts)

	if err != nil {
		// The panic built-in function stops normal execution of the current goroutine.
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}
