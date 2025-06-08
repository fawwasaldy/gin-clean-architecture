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

	userSchema := EntityToSchema(userEntity)
	if err = db.WithContext(ctx).Create(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity = SchemaToEntity(userSchema)
	return userEntity, nil
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

	var userSchema Schema
	if err = db.WithContext(ctx).Where("id = ?", id).Take(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity := SchemaToEntity(userSchema)
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

	var userSchema Schema
	if err = db.WithContext(ctx).Where("email = ?", email).Take(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity := SchemaToEntity(userSchema)
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

	var userSchema Schema
	if err = db.WithContext(ctx).Where("email = ?", email).Take(&userSchema).Error; err != nil {
		return user.User{}, false, err
	}

	userEntity := SchemaToEntity(userSchema)
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

	userSchema := EntityToSchema(userEntity)
	if err = db.WithContext(ctx).Updates(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity = SchemaToEntity(userSchema)
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

	if err = db.WithContext(ctx).Where("id = ?", id).Delete(&Schema{}).Error; err != nil {
		return err
	}

	return nil
}
