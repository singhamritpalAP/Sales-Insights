package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sales/internal/constants"
	"sales/internal/services"
	"sales/internal/utils"
)

// RefreshHandler handles the data refresh endpoint.
func RefreshHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := services.RefreshDatabase(db)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.String(http.StatusOK, "Data refreshed successfully.")
	}
}

// GetTopProductsOverallHandler handles the retrieval of top N products overall.
func GetTopProductsOverallHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nStr := ctx.Query(constants.Limit)
		startDate := ctx.Query(constants.StartDate)
		endDate := ctx.Query(constants.EndDate)

		n, err := utils.ValidateParamsAndGetLimit(nStr, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		topProducts, err := services.GetTopProductsOverall(db, n, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, topProducts)
	}
}

// GetTopProductsByCategoryHandler handles the retrieval of top N products by category.
func GetTopProductsByCategoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nStr := ctx.Query("n")
		startDate := ctx.Query("start_date")
		endDate := ctx.Query("end_date")

		n, err := utils.ValidateParamsAndGetLimit(nStr, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		topProductsByCategory, err := services.GetTopProductsByCategory(db, n, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, topProductsByCategory)
	}
}

// GetTopProductsByRegionHandler handles the retrieval of top N products by region.
func GetTopProductsByRegionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		nStr := ctx.Query("n")
		startDate := ctx.Query("start_date")
		endDate := ctx.Query("end_date")

		n, err := utils.ValidateParamsAndGetLimit(nStr, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		topProductsByRegion, err := services.GetTopProductsByRegion(db, n, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, topProductsByRegion)
	}
}
