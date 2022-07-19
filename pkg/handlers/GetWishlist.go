package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melvin2016/wishlist-service/pkg/mocks"
)

// GetWishlist gets the items from wishlist and convert it into json
func GetWishlist(c *gin.Context) {
	c.JSON(http.StatusOK, mocks.Wishlists)
}
