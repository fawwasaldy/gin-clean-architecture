package user

import (
	"context"
	"kpl-base/platform/pagination"
)

type (
	Repository interface {
		Register(ctx context.Context, tx interface{}, userEntity User) (User, error)
		GetAllUsersWithPagination(
			ctx context.Context,
			tx interface{},
			req pagination.Request,
		) (pagination.ResponseWithData, error)
		GetUserByID(ctx context.Context, tx interface{}, id string) (User, error)
		GetUserByEmail(ctx context.Context, tx interface{}, email string) (User, error)
		CheckEmail(ctx context.Context, tx interface{}, email string) (User, bool, error)
		Update(ctx context.Context, tx interface{}, userEntity User) (User, error)
		Delete(ctx context.Context, tx interface{}, id string) error
	}
)
