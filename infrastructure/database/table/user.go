package table

import (
	"time"

	"github.com/fawwasaldy/gin-clean-architecture/domain/identity"
	"github.com/fawwasaldy/gin-clean-architecture/domain/shared"
	"github.com/fawwasaldy/gin-clean-architecture/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	Name        string         `gorm:"type:varchar(100);not null;column:name"`
	Email       string         `gorm:"type:varchar(255);uniqueIndex;not null;column:email"`
	PhoneNumber string         `gorm:"type:varchar(20);index;column:phone_number"`
	Password    string         `gorm:"type:varchar(255);not null;column:password"`
	Role        string         `gorm:"type:varchar(50);not null;default:'user';column:role"`
	ImageUrl    string         `gorm:"type:varchar(255);column:image_url"`
	IsVerified  bool           `gorm:"default:false;column:is_verified"`
	CreatedAt   time.Time      `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt   time.Time      `gorm:"type:timestamp with time zone;column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp with time zone;column:deleted_at"`
}

func UserEntityToTable(entity user.User) User {
	var deletedAtTime time.Time
	if entity.Timestamp.DeletedAt != nil {
		deletedAtTime = *entity.Timestamp.DeletedAt
	} else {
		deletedAtTime = time.Time{}
	}
	return User{
		ID:          entity.ID.ID,
		Name:        entity.Name,
		Email:       entity.Email,
		PhoneNumber: entity.PhoneNumber,
		Password:    entity.Password.Password,
		Role:        entity.Role.Name,
		ImageUrl:    entity.ImageUrl.Path,
		IsVerified:  entity.IsVerified,
		CreatedAt:   entity.Timestamp.CreatedAt,
		UpdatedAt:   entity.Timestamp.UpdatedAt,
		DeletedAt: gorm.DeletedAt{
			Time:  deletedAtTime,
			Valid: entity.Timestamp.DeletedAt != nil,
		},
	}
}

func UserTableToEntity(table User) user.User {
	return user.User{
		ID:          identity.NewIDFromTable(table.ID),
		Name:        table.Name,
		Email:       table.Email,
		PhoneNumber: table.PhoneNumber,
		Password:    user.NewPasswordFromTable(table.Password),
		Role:        user.NewRoleFromTable(table.Role),
		ImageUrl:    shared.NewURLFromTable(table.ImageUrl),
		IsVerified:  table.IsVerified,
		Timestamp: shared.Timestamp{
			CreatedAt: table.CreatedAt,
			UpdatedAt: table.UpdatedAt,
			DeletedAt: &table.DeletedAt.Time,
		},
	}
}
