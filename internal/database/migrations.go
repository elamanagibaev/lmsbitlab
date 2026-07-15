package database

import (
	"LMSBitLab/migrations"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	goose.SetBaseFS(migrations.FS)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(sqlDB, "."); err != nil {
		return err
	}

	return nil
}
