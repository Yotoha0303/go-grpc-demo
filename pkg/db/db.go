package db

import (
	"fmt"

	"go-grpc-demo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init opens a GORM connection with the given MySQL config.
func Init(cfg *config.MySQLConfig) (*gorm.DB, error) {
	if cfg == nil {
		return nil, fmt.Errorf("mysql config is nil")
	}

	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql.DB: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping mysql: %w", err)
	}

	return db, nil
}
