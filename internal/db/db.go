package db

import (
	"fmt"

	"github.com/Shyyw1e/effective-mobile-subs/pkg/config"
	"github.com/Shyyw1e/effective-mobile-subs/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logger.Log.Errorf("failed to open database: %v", err)
		return nil, err
	}

	if err := database.AutoMigrate(&Subscription{}); err != nil {
		logger.Log.Errorf("failed to migrate: %v", err)
		return nil, err
	}

	logger.Log.Info("Migration successful")
	return database, nil
}