package database

import (
	"gorm.io/gorm"
)

// DB represents the connection to the database
var (
	DB *gorm.DB
)