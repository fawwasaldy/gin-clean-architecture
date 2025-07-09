package user

import (
	"gin-clean-architecture/domain/identity"
	"gin-clean-architecture/domain/shared"
)

type User struct {
	ID          identity.ID
	Name        string
	Email       string
	PhoneNumber string
	Password    Password
	Role        Role
	ImageUrl    shared.URL
	IsVerified  bool
	shared.Timestamp
}
