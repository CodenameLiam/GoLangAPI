package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// Book is the structure used to represent books within the application
type Book struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;" json:"ID"`
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int8 `json:"rating"`
	PublishedDate time.Time `json:"publishedDate"`
}