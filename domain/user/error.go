package user

import "errors"

var (
	ErrorCreateUser         = errors.New("failed to create user")
	ErrorGetAllUsers        = errors.New("failed to get all users")
	ErrorGetUserById        = errors.New("failed to get user by id")
	ErrorGetUserByEmail     = errors.New("failed to get user by email")
	ErrorEmailAlreadyExists = errors.New("email already exist")
	ErrorUpdateUser         = errors.New("failed to update user")
	ErrorUserNotFound       = errors.New("user not found")
	ErrorEmailNotFound      = errors.New("email not found")
	ErrorDeleteUser         = errors.New("failed to delete user")
	ErrorTokenInvalid       = errors.New("token invalid")
	ErrorTokenExpired       = errors.New("token expired")
)
