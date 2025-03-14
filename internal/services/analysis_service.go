package services

import (
	"sales/internal/models"
	"sales/internal/repository"

	"gorm.io/gorm"
)

func GetTopProductsOverall(db *gorm.DB, n int, startDate string, endDate string) ([]models.Product, error) {
	return repository.GetTopProductsOverall(db, n, startDate, endDate)
}

func GetTopProductsByCategory(db *gorm.DB, n int, startDate string, endDate string) (map[string][]models.Product, error) {
	return repository.GetTopProductsByCategory(db, n, startDate, endDate)
}

func GetTopProductsByRegion(db *gorm.DB, n int, startDate string, endDate string) (map[string][]models.Product, error) {
	return repository.GetTopProductsByRegion(db, n, startDate, endDate)
}
