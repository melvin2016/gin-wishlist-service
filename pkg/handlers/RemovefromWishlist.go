package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/melvin2016/wishlist-service/pkg/mocks"
)

// RemoveFromWishlist removes an item from the wishlist
func RemoveFromWishlist(c *gin.Context) {
	// get the "id" from the request
	params := c.Params
	idstr, idFound := params.Get("id")
	if !idFound {
		return
	}

	// convert id from string to int
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return
	}

	// check if the id is present
	var idx int
	for indx, val := range mocks.Wishlists {
		if val.ID == id {
			idx = indx
		}
	}

	// remove the id specified from the array
	mocks.Wishlists = append(mocks.Wishlists[:idx], mocks.Wishlists[idx+1:]...)

	c.JSON(http.StatusOK, mocks.Wishlists)
}
