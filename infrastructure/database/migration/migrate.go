package migration

import (
	"gin-clean-architecture/infrastructure/database/refresh_token"
	"gin-clean-architecture/infrastructure/database/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&user.User{},
		&refresh_token.RefreshToken{},
	); err != nil {
		return err
	}

	return nil
}
