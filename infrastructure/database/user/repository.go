package user

import (
	"context"
	"kpl-base/domain/user"
	"kpl-base/infrastructure/database/transaction"
	"kpl-base/infrastructure/database/validation"
	"kpl-base/platform/pagination"
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

func (r *repository) GetAllUsersWithPagination(ctx context.Context, tx interface{}, req pagination.Request) (pagination.ResponseWithData, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return pagination.ResponseWithData{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchemas []User
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&User{})
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
		data[i] = SchemaToEntity(userSchema)
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

func (r *repository) GetUserByID(ctx context.Context, tx interface{}, id string) (user.User, error) {
	validatedTransaction, err := validation.ValidateTransaction(tx)
	if err != nil {
		return user.User{}, err
	}

	db := validatedTransaction.DB()
	if db == nil {
		db = r.db.DB()
	}

	var userSchema User
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

	var userSchema User
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

	var userSchema User
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

	if err = db.WithContext(ctx).Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}
