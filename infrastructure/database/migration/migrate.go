package migration

import (
	"gin-clean-architecture/infrastructure/database/schema"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&schema.User{},
		&schema.RefreshToken{},
	); err != nil {
		return err
	}

	return nil
}
