package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/melvin2016/wishlist-service/pkg/models"
)

// RemoveFromWishlist removes an item from the wishlist
func (h Handler) RemoveFromWishlist(c *gin.Context) {

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

	// find the wishlist that needs to be deleted
	toDeleteWishlist := &models.Wishlist{}
	if result := h.DB.Find(toDeleteWishlist, id); result.Error != nil {
		log.Println(result.Error)
	}

	// delete wishlist matching id
	deletedWishlist := models.Wishlist{}
	if result := h.DB.Delete(&deletedWishlist, id); result.Error != nil {
		log.Println(result.Error)
	}

	// send response
	c.JSON(http.StatusOK, toDeleteWishlist)
}
