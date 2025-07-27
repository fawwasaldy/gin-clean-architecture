package command

import (
	"gin-clean-architecture/infrastructure/database/migration"
	"log"
	"os"

	"gorm.io/gorm"
)

func Commands(db *gorm.DB) bool {
	migrate := false
	seed := false
	rollback := false
	run := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
		}
		if arg == "--rollback" {
			rollback = true
		}
		if arg == "--run" {
			run = true
		}
	}

	if migrate {
		if err := migration.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}

	if seed {
		if err := migration.Seeder(db); err != nil {
			log.Fatalf("error migration seeder: %v", err)
		}
		log.Println("seeder completed successfully")
	}

	if rollback {
		if err := migration.Rollback(db); err != nil {
			log.Fatalf("error migration rollback: %v", err)
		}
		log.Println("rollback completed successfully")
	}

	if run {
		return true
	}

	return false
}
