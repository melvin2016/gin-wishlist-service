package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/melvin2016/wishlist-service/pkg/db"
	"github.com/melvin2016/wishlist-service/pkg/handlers"
)

func main() {
	log.SetPrefix("server startup: ") // sets prefix for all the logs in this files
	log.SetFlags(0)                   // remove time and date

	// connect database
	db := db.Init()

	// initialize route handlers
	h := handlers.New(db)

	// get the router
	router := gin.Default()

	// start db connection

	// gets all the current wishlist items
	router.GET("/wishlist", h.GetWishlist)

	// adds an item to the wishlist
	router.POST("/wishlist", h.AddToWishlist)

	// removes an item from the wishlist
	router.DELETE("/wishlist/:id", h.RemoveFromWishlist)

	err := router.Run("localhost:8080") // starts server on localhost:8080
	if err != nil {
		log.Fatal(err) // if any error on statup log to console
	}
}
