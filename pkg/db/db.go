package db

import (
	"log"

	"github.com/melvin2016/wishlist-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init starts the postgres database and returns a connection
func Init() *gorm.DB {
	dsn := "postgres://admin:admin1234@localhost:5432/wishlist"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Wishlist{})
	return db
}
