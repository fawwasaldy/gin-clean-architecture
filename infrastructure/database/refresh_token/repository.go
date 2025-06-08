package refresh_token

import (
	"context"
	"kpl-base/domain/refresh_token"
	"kpl-base/infrastructure/database/transaction"
	"kpl-base/infrastructure/database/validation"
	"time"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) refresh_token.Repository {
	return &repository{db: db}
}

func (r repository) Create(ctx context.Context, tx interface{}, refreshTokenEntity refresh_token.RefreshToken) (refresh_token.RefreshToken, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return refresh_token.RefreshToken{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err := db.WithContext(ctx).Create(&refreshTokenEntity).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	return refreshTokenEntity, nil
}

func (r repository) FindByToken(ctx context.Context, tx interface{}, token string) (refresh_token.RefreshToken, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return refresh_token.RefreshToken{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var refreshTokenEntity refresh_token.RefreshToken
	if err := db.WithContext(ctx).Where("token = ?", token).Take(&refreshTokenEntity).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	return refreshTokenEntity, nil
}

func (r repository) DeleteByUserID(ctx context.Context, tx interface{}, userID string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err := db.WithContext(ctx).Where("user_id = ?", userID).Delete(&refresh_token.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteByToken(ctx context.Context, tx interface{}, token string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err := db.WithContext(ctx).Where("token = ?", token).Delete(&refresh_token.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) DeleteExpired(ctx context.Context, tx interface{}) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err := db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&refresh_token.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}
