package command

import (
	"kpl-base/infrastructure/database/migration"
	"log"
	"os"

	"gorm.io/gorm"
)

func Commands(db *gorm.DB) bool {
	migrate := false
	seed := false
	run := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
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
		//if err := migration.Seeder(db); err != nil {
		//	log.Fatalf("error migration seeder: %v", err)
		//}
		//log.Println("seeder completed successfully")
	}

	if run {
		return true
	}

	return false
}
