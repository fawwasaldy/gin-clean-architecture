package seed

import (
	"gin-clean-architecture/domain/shared"
	"gin-clean-architecture/domain/user"
	"gin-clean-architecture/infrastructure/database/migration/data"
	"gin-clean-architecture/infrastructure/database/table"
	"gorm.io/gorm"
	"log"
)

func User(db *gorm.DB) error {
	for _, userData := range data.UserSeedData() {
		var existingUser table.User
		if err := db.Where("email = ?", userData.Email).First(&existingUser).Error; err == nil {
			log.Printf("User with email %s already exists, skipping seed process for this user.", userData.Email)
			continue
		}

		password, err := user.NewPassword(userData.Password)
		if err != nil {
			log.Fatalf("Failed to create password for seeder: %v", err)
			return err
		}

		role, err := user.NewRole(userData.Role)
		if err != nil {
			log.Fatalf("Failed to create role for seeder: %v", err)
			return err
		}

		imageUrl, err := shared.NewURL(userData.ImageUrl)
		if err != nil {
			log.Fatalf("Failed to create image URL for seeder: %v", err)
			return err
		}

		userEntity := user.User{
			Name:        userData.Name,
			Email:       userData.Email,
			PhoneNumber: userData.PhoneNumber,
			Password:    password,
			Role:        role,
			ImageUrl:    imageUrl,
			IsVerified:  userData.IsVerified,
		}

		userTable := table.UserEntityToTable(userEntity)
		if err = db.Create(&userTable).Error; err != nil {
			log.Fatalf("Failed to run user seeder: %v", err)
			return err
		}
	}
	return nil
}
