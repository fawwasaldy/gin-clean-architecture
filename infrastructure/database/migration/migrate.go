package migration

import (
	"gorm.io/gorm"
	"kpl-base/infrastructure/database/refresh_token"
	"kpl-base/infrastructure/database/user"
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
