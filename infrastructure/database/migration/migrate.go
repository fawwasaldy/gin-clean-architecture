package migration

import (
	"fmt"
	"gin-clean-architecture/infrastructure/database/config"
	"gin-clean-architecture/infrastructure/database/table"
	"os"

	"gorm.io/gorm"
)

var entities = []interface{}{
	&table.User{},
	&table.RefreshToken{},
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(entities...); err != nil {
		return err
	}

	return nil
}

func Rollback(db *gorm.DB) error {
	if os.Getenv("APP_ENV") == config.RunProduction {
		return fmt.Errorf("rollback is not allowed for production environment")
	}

	if err := db.Migrator().DropTable(entities...); err != nil {
		return err
	}

	return nil
}
