package models

import "time"

// CatalogType Enums
type CatalogType int

const (
	ITEM CatalogType = iota
	BUNDLE
	COMPOSITE_ITEM
)

// CatalogCondition Enums
type CatalogCondition int

const (
	BRAND_NEW CatalogCondition = iota
	REFURBISHED
)

// Wishlist structure
type Wishlist struct {
	ID                  int              `json:"id" gorm:"primaryKey"`
	UserID              int              `json:"user_id"`
	CatalogID           int              `json:"catalog_id"`
	CatalogName         string           `json:"catalog_name"`
	CatalogType         CatalogType      `json:"catalog_type"`
	CatalogImageURL     string           `json:"catalog_image_url"`
	CatalogCondition    CatalogCondition `json:"catalog_condition"`
	CatalogStrikePrice  float64          `json:"catalog_strike_price"`
	CatalogSellingPrice float64          `json:"catalog_selling_price"`
	CreatedAt           time.Time        `json:"created_at"`
	DeletedAt           time.Time        `json:"deleted_at"`
}
