package database

import (
	"gorm.io/gorm"
	"sales/internal/models"
)

// AutoMigrateSchemas automatically migrates the database schemas.
func AutoMigrateSchemas(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Product{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		return err
	}
	return nil
}
