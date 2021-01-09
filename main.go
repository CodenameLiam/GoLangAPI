package main

import(
	"fmt"
	"github.com/CodenameLiam/GoLangAPI/api"
	"github.com/CodenameLiam/GoLangAPI/models"
	"github.com/CodenameLiam/GoLangAPI/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialise database
	initialiseDatabase()

	// Run API
	api.Run()

} 

func initialiseDatabase() {
	fmt.Println("Test")

	dsn := "host=localhost user=postgres password=pass dbname=Books port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	database.DB = db
	fmt.Println("Database connection successfully opened")

	database.DB.AutoMigrate(&models.Book{})
	fmt.Println("Database migration successfully completed")
}

