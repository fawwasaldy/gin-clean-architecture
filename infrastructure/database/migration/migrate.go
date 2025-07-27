package migration

import (
	"gin-clean-architecture/infrastructure/database/table"
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
	if err := db.Migrator().DropTable(entities...); err != nil {
		return err
	}

	return nil
}
