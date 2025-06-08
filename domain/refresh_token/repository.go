package refresh_token

import (
	"context"
)

type (
	Repository interface {
		Create(ctx context.Context, tx interface{}, refreshTokenEntity RefreshToken) (RefreshToken, error)
		FindByToken(ctx context.Context, tx interface{}, token string) (RefreshToken, error)
		DeleteByUserID(ctx context.Context, tx interface{}, userID string) error
		DeleteByToken(ctx context.Context, tx interface{}, token string) error
		DeleteExpired(ctx context.Context, tx interface{}) error
	}
)
