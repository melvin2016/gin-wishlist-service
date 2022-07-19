package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melvin2016/wishlist-service/pkg/mocks"
	"github.com/melvin2016/wishlist-service/pkg/models"
)

// AddToWishlist adds an item to the wishlist array
func AddToWishlist(c *gin.Context) {
	var newWishlist models.Wishlist

	if err := c.BindJSON(&newWishlist); err != nil {
		log.Panic(err)
		return
	}

	mocks.Wishlists = append(mocks.Wishlists, newWishlist)

	c.IndentedJSON(http.StatusCreated, newWishlist)
}
