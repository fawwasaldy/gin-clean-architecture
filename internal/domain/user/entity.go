package user

import (
	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/identity"
	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/shared"
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
