package repository

import (
	"context"
	"time"

	"github.com/fawwasaldy/gin-clean-architecture/domain/refresh_token"
	"github.com/fawwasaldy/gin-clean-architecture/infrastructure/database/table"
	"github.com/fawwasaldy/gin-clean-architecture/infrastructure/database/transaction"
	"github.com/fawwasaldy/gin-clean-architecture/infrastructure/database/validation"
)

type refreshTokenRepository struct {
	db *transaction.Repository
}

func NewRefreshTokenRepository(db *transaction.Repository) refresh_token.Repository {
	return &refreshTokenRepository{db: db}
}

func (r refreshTokenRepository) Create(ctx context.Context, tx interface{}, refreshTokenEntity refresh_token.RefreshToken) (refresh_token.RefreshToken, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return refresh_token.RefreshToken{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	refreshTokenTable := table.RefreshTokenEntityToTable(refreshTokenEntity)
	if err = db.WithContext(ctx).Create(&refreshTokenTable).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	refreshTokenEntity = table.RefreshTokenTableToEntity(refreshTokenTable)
	return refreshTokenEntity, nil
}

func (r refreshTokenRepository) FindByUserID(ctx context.Context, tx interface{}, userID string) (refresh_token.RefreshToken, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return refresh_token.RefreshToken{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var refreshTokenTable table.RefreshToken
	if err = db.WithContext(ctx).Where("user_id = ?", userID).Take(&refreshTokenTable).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	refreshTokenEntity := table.RefreshTokenTableToEntity(refreshTokenTable)
	return refreshTokenEntity, nil
}

func (r refreshTokenRepository) DeleteByUserID(ctx context.Context, tx interface{}, userID string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err = db.WithContext(ctx).Where("user_id = ?", userID).Delete(&table.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r refreshTokenRepository) DeleteByToken(ctx context.Context, tx interface{}, token string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err = db.WithContext(ctx).Where("token = ?", token).Delete(&table.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r refreshTokenRepository) DeleteExpired(ctx context.Context, tx interface{}) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err = db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&table.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}
