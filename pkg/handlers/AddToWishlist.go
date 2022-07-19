package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melvin2016/wishlist-service/pkg/models"
)

// AddToWishlist adds an item to the wishlist array
func (h Handler) AddToWishlist(c *gin.Context) {
	var newWishlist models.Wishlist

	// bind the data to variable
	err := c.BindJSON(&newWishlist)
	if err != nil {
		log.Panic(err)
		return
	}

	// create a new wishlist
	result := h.DB.Create(&newWishlist)
	if result.Error != nil {
		log.Println(result.Error)
	}

	c.IndentedJSON(http.StatusCreated, newWishlist)
}
