package handlers

import "gorm.io/gorm"

// Handler refers to the object containing the db connection
type Handler struct {
	DB *gorm.DB
}

// New makes a handler with db atached to it
func New(db *gorm.DB) Handler {
	return Handler{
		DB: db,
	}
}
