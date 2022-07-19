package mocks

import (
	"time"

	"github.com/melvin2016/wishlist-service/pkg/models"
)

// Wishlists contains mocks data
var Wishlists = []models.Wishlist{{
	ID:                  1,
	UserID:              1,
	CatalogID:           1,
	CatalogName:         "furniture",
	CatalogType:         models.ITEM,
	CatalogImageURL:     "https://adad.com",
	CatalogCondition:    models.BRAND_NEW,
	CatalogStrikePrice:  1200.34,
	CatalogSellingPrice: 1000.00,
	CreatedAt:           time.Now(),
	DeletedAt:           time.Now(),
}}
