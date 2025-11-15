package migration

import (
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/migration/seed"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seed.User(db); err != nil {
		return err
	}

	return nil
}
