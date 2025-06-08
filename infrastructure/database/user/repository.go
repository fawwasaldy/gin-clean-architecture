package user

import (
	"context"
	"kpl-base/domain/user"
	"kpl-base/infrastructure/database/transaction"
	"kpl-base/infrastructure/database/validation"
)

type repository struct {
	db *transaction.Repository
}

func NewRepository(db *transaction.Repository) user.Repository {
	return &repository{db: db}
}

func (r *repository) Register(ctx context.Context, tx interface{}, userEntity user.User) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err := db.WithContext(ctx).Create(&userEntity).Error; err != nil {
		return user.User{}, err
	}

	return user.User{}, nil
}

func (r *repository) GetUserByID(ctx context.Context, tx interface{}, id string) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userEntity user.User
	if err := db.WithContext(ctx).Where("id = ?", id).Take(&userEntity).Error; err != nil {
		return user.User{}, err
	}

	return userEntity, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, tx interface{}, email string) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userEntity user.User
	if err := db.WithContext(ctx).Where("email = ?", email).Take(&userEntity).Error; err != nil {
		return user.User{}, err
	}

	return userEntity, nil
}

func (r *repository) CheckEmail(ctx context.Context, tx interface{}, email string) (user.User, bool, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, false, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userEntity user.User
	if err := db.WithContext(ctx).Where("email = ?", email).Take(&userEntity).Error; err != nil {
		return user.User{}, false, err
	}

	return userEntity, true, nil
}

func (r *repository) Update(ctx context.Context, tx interface{}, userEntity user.User) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err := db.WithContext(ctx).Updates(&userEntity).Error; err != nil {
		return user.User{}, err
	}

	return userEntity, nil
}

func (r *repository) Delete(ctx context.Context, tx interface{}, id string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err := db.WithContext(ctx).Where("id = ?", id).Delete(&user.User{}).Error; err != nil {
		return err
	}

	return nil
}
