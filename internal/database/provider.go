package database

import (
	"agnos-assignment/internal/config"
	"agnos-assignment/internal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideConnection(cfg *config.Config) (*Connection, error) {
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to configure connection pool: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)

	if err := db.AutoMigrate(&models.Staff{}, &models.Patient{}); err != nil {
		return nil, fmt.Errorf("failed to run auto-migrations: %w", err)
	}

	if err := RunSQLMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run SQL migrations: %w", err)
	}

	return &Connection{DB: db}, nil
}
