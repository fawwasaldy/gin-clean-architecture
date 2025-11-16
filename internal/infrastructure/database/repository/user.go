package repository

import (
	"context"

	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/user"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/table"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/transaction"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/validation"
	"github.com/fawwasaldy/gin-clean-architecture/platform/pagination"
	"github.com/samber/do/v2"
)

type userRepository struct {
	db *transaction.Repository
}

func NewUserRepository(injector do.Injector) user.Repository {
	db := do.MustInvoke[*transaction.Repository](injector)
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

	userTable := table.UserEntityToTable(userEntity)
	if err = db.WithContext(ctx).Create(&userTable).Error; err != nil {
		return user.User{}, err
	}

	userEntity = table.UserTableToEntity(userTable)
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

	var userTables []table.User
	var count int64

	req.Default()

	query := db.WithContext(ctx).Model(&table.User{})
	if req.Search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	if err = query.Count(&count).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	if err = query.Scopes(pagination.Paginate(req)).Find(&userTables).Error; err != nil {
		return pagination.ResponseWithData{}, err
	}

	totalPage := pagination.TotalPage(count, int64(req.PerPage))

	data := make([]any, len(userTables))
	for i, userTable := range userTables {
		data[i] = table.UserTableToEntity(userTable)
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

	var userTable table.User
	if err = db.WithContext(ctx).Where("id = ?", id).Take(&userTable).Error; err != nil {
		return user.User{}, err
	}

	userEntity := table.UserTableToEntity(userTable)
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

	var userTable table.User
	if err = db.WithContext(ctx).Where("email = ?", email).Take(&userTable).Error; err != nil {
		return user.User{}, err
	}

	userEntity := table.UserTableToEntity(userTable)
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

	var userTable table.User
	if err = db.WithContext(ctx).Where("email = ?", email).Take(&userTable).Error; err != nil {
		return user.User{}, false, err
	}

	userEntity := table.UserTableToEntity(userTable)
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

	userTable := table.UserEntityToTable(userEntity)
	if err = db.WithContext(ctx).Updates(&userTable).Error; err != nil {
		return user.User{}, err
	}

	userEntity = table.UserTableToEntity(userTable)
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

	if err = db.WithContext(ctx).Where("id = ?", id).Delete(&table.User{}).Error; err != nil {
		return err
	}

	return nil
}
