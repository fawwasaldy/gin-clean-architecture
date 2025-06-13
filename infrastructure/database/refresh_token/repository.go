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

	refreshTokenSchema := EntityToSchema(refreshTokenEntity)
	if err = db.WithContext(ctx).Create(&refreshTokenSchema).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	refreshTokenEntity = SchemaToEntity(refreshTokenSchema)
	return refreshTokenEntity, nil
}

func (r repository) FindByUserID(ctx context.Context, tx interface{}, userID string) (refresh_token.RefreshToken, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return refresh_token.RefreshToken{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var refreshTokenSchema RefreshToken
	if err = db.WithContext(ctx).Where("user_id = ?", userID).Take(&refreshTokenSchema).Error; err != nil {
		return refresh_token.RefreshToken{}, err
	}

	refreshTokenEntity := SchemaToEntity(refreshTokenSchema)
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

	if err = db.WithContext(ctx).Where("user_id = ?", userID).Delete(&RefreshToken{}).Error; err != nil {
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

	if err = db.WithContext(ctx).Where("token = ?", token).Delete(&RefreshToken{}).Error; err != nil {
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

	if err = db.WithContext(ctx).Where("expires_at < ?", time.Now()).Delete(&RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}
