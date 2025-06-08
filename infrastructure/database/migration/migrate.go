package migrations

import (
	"gorm.io/gorm"
	"kpl-base/domain/refresh_token"
	"kpl-base/domain/user"
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
