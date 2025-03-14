package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // Blank import to register the driver
)

func NewDatabase(dbPath string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm database: %w", err)
	}
	err = AutoMigrateSchemas(gormDB)
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
