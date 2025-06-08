package user

import "fmt"

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

var (
	Roles = []Role{
		{RoleAdmin},
		{RoleUser},
	}
)

type Role struct {
	name string
}

func NewRole(name string) (Role, error) {
	if !isValidRole(name) {
		return Role{}, fmt.Errorf("invalid role name")
	}
	return Role{
		name: name,
	}, nil
}

func (r Role) Name() string {
	return r.name
}

func isValidRole(name string) bool {
	for _, role := range Roles {
		if role.Name() == name {
			return true
		}
	}
	return false
}
