package repository

import (
	"context"
	"gin-clean-architecture/domain/refresh_token"
	"gin-clean-architecture/infrastructure/database/schema"
	"gin-clean-architecture/infrastructure/database/transaction"
	"gin-clean-architecture/infrastructure/database/validation"
	"time"
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

	refreshTokenSchema := schema.RefreshTokenEntityToSchema(refreshTokenEntity)
	if err = db.WithContext(ctx).Create(&refreshTokenSchema).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	refreshTokenEntity = schema.RefreshTokenSchemaToEntity(refreshTokenSchema)
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

	var refreshTokenSchema schema.RefreshToken
	if err = db.WithContext(ctx).Where("user_id = ?", userID).Take(&refreshTokenSchema).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	refreshTokenEntity := schema.RefreshTokenSchemaToEntity(refreshTokenSchema)
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

	if err = db.WithContext(ctx).Where("user_id = ?", userID).Delete(&schema.RefreshToken{}).Error; err != nil {
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

	if err = db.WithContext(ctx).Where("token = ?", token).Delete(&schema.RefreshToken{}).Error; err != nil {
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

	if err = db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&schema.RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}
