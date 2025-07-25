package repository

import (
	"context"
	"gin-clean-architecture/domain/user"
	"gin-clean-architecture/infrastructure/database/schema"
	"gin-clean-architecture/infrastructure/database/transaction"
	"gin-clean-architecture/infrastructure/database/validation"
	"gin-clean-architecture/platform/pagination"
)

type userRepository struct {
	db *transaction.Repository
}

func NewUserRepository(db *transaction.Repository) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Register(ctx context.Context, tx interface{}, userEntity user.User) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	userSchema := schema.UserEntityToSchema(userEntity)
	if err = db.WithContext(ctx).Create(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity = schema.UserSchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *userRepository) GetAllUsersWithPagination(ctx context.Context, tx interface{}, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchemas []schema.User
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&schema.User{})
	if req.Search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&userSchemas).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	data := make([]any, len(userSchemas))
	for i, userSchema := range userSchemas {
		data[i] = schema.UserSchemaToEntity(userSchema)
	}
	return pagination.ResponseWithData{
		Data: data,
		Response: pagination.Response{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *userRepository) GetUserByID(ctx context.Context, tx interface{}, id string) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchema schema.User
	if err = db.WithContext(ctx).Where("id = ?", id).Take(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity := schema.UserSchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, tx interface{}, email string) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchema schema.User
	if err = db.WithContext(ctx).Where("email = ?", email).Take(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity := schema.UserSchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *userRepository) CheckEmail(ctx context.Context, tx interface{}, email string) (user.User, bool, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, false, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchema schema.User
	if err = db.WithContext(ctx).Where("email = ?", email).Take(&userSchema).Error; err != nil {
		return user.User{}, false, err
	}

	userEntity := schema.UserSchemaToEntity(userSchema)
	return userEntity, true, nil
}

func (r *userRepository) Update(ctx context.Context, tx interface{}, userEntity user.User) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	userSchema := schema.UserEntityToSchema(userEntity)
	if err = db.WithContext(ctx).Updates(&userSchema).Error; err != nil {
		return user.User{}, err
	}

	userEntity = schema.UserSchemaToEntity(userSchema)
	return userEntity, nil
}

func (r *userRepository) Delete(ctx context.Context, tx interface{}, id string) error {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	if err = db.WithContext(ctx).Where("id = ?", id).Delete(&schema.User{}).Error; err != nil {
		return err
	}

	return nil
}
