package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melvin2016/wishlist-service/pkg/models"
)

// GetWishlist gets the items from wishlist and convert it into json
func (h Handler) GetWishlist(c *gin.Context) {
	var wishlists []models.Wishlist

	// get all the wishlists
	result := h.DB.Find(&wishlists)
	if result.Error != nil {
		log.Println(result.Error)
	}

	// sent wishlists response
	c.JSON(http.StatusOK, wishlists)
}
