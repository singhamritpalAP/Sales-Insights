package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes initializes the routes for the application.
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/refresh", RefreshHandler(db))
	router.GET("/top-products/overall", GetTopProductsOverallHandler(db))
	router.GET("/top-products/category", GetTopProductsByCategoryHandler(db))
	router.GET("/top-products/region", GetTopProductsByRegionHandler(db))
}
