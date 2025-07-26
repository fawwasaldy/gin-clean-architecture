package data

import "gin-clean-architecture/domain/user"

type UserSeed struct {
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	Role        string
	ImageUrl    string
	IsVerified  bool
}

func UserSeedData() []UserSeed {
	return []UserSeed{
		{
			Name:        "Dummy Admin",
			Email:       "admin@example.com",
			PhoneNumber: "1234567890",
			Password:    "adminpassword123",
			Role:        user.RoleAdmin,
			ImageUrl:    "profile/default.png",
			IsVerified:  false,
		},
		{
			Name:        "Dummy User 1",
			Email:       "user1@example.com",
			PhoneNumber: "12345678901",
			Password:    "userpassword123",
			Role:        user.RoleUser,
			ImageUrl:    "profile/default.png",
			IsVerified:  false,
		},
		{
			Name:        "Dummy User 2",
			Email:       "user2@example.com",
			PhoneNumber: "12345678902",
			Password:    "userpassword123",
			Role:        user.RoleUser,
			ImageUrl:    "profile/default.png",
			IsVerified:  false,
		},
		{
			Name:        "Dummy User 3",
			Email:       "user3@example.com",
			PhoneNumber: "12345678903",
			Password:    "userpassword123",
			Role:        user.RoleUser,
			ImageUrl:    "profile/default.png",
			IsVerified:  false,
		},
	}
}
