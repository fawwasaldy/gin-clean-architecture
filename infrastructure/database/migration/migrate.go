package migration

import (
	"gin-clean-architecture/infrastructure/database/table"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&table.User{},
		&table.RefreshToken{},
	); err != nil {
		return err
	}

	return nil
}
