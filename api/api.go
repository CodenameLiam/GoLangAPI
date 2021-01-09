package api

import (
	"log"
	"time"
	"errors"
	"gorm.io/gorm"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2"
	"github.com/CodenameLiam/GoLangAPI/models"
	"github.com/CodenameLiam/GoLangAPI/database"
)



// Run starts the API
func Run() {
	// Initialise server
	app := fiber.New()

	// Define routes
	app.Get("/api/books/:id", getBook)
	app.Get("/api/books", getBooks)
	app.Post("/api/books", createBook)
	app.Put("/api/books/:id", updateBook)
	app.Delete("/api/books/:id", deleteBook)

	// Servercontent	
	log.Fatal(app.Listen(":3000"))
}

func getBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []models.Book
	db.Find(&books)
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book models.Book
	db.First(&book, id)
	return c.JSON(book)
}

func createBook(c *fiber.Ctx) error {
	db := database.DB
	book := new(models.Book)

	var body map[string]interface{}
	json.Unmarshal(c.Body(), &body)

	if err := c.BodyParser(book); err != nil {
		c.Status(503).JSON(err)
	}

	book.ID = uuid.New()

	if book.Title == "" {
		return c.Status(500).SendString("Please enter a valid title")
	}
	if book.Author == "" {
		return c.Status(500).SendString("Please enter a valid Author")
	}
	if book.PublishedDate.IsZero() {
		return c.Status(500).SendString("Please enter a valid Date")
	}

	publishedDate, err := time.Parse("2006-01-02", "1982-02-23")
	if err != nil {
		c.Status(503).JSON(err)
	}
	book.PublishedDate = publishedDate

	db.Create(&book)
	return c.JSON(book)


	// var book models.Book





	// time, err := time.Parse("2006-01-02", "1982-02-23")

	// if err != nil {
	// 	panic("Time is messed")
	// }

	// book.ID = uuid.New()
	// book.Title = "Dune"
	// book.Author = "Robert"
	// book.Rating = 10
	// book.PublishedDate = time

	// db.Create(&book)
	// return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	return c.SendString("Update book")
}

func deleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book models.Book

	parseID, errID := uuid.Parse(id)

	if errID != nil {
		return c.Status(500).SendString("Invalid input")
	}

	err := db.First(&book, parseID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).SendString("Record not found")
	}
	db.Delete(&book)
	return c.SendString("Book successfully deleted")
	
}