package user

import (
	"github.com/google/uuid"
	"kpl-base/domain/identity"
	"kpl-base/domain/shared"
	"kpl-base/domain/user"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4();column:id"`
	Name        string    `gorm:"type:varchar(100);not null;column:name"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null;column:email"`
	PhoneNumber string    `gorm:"type:varchar(20);index;column:phone_number"`
	Password    string    `gorm:"type:varchar(255);not null;column:password"`
	Role        string    `gorm:"type:varchar(50);not null;default:'user';column:role"`
	ImageUrl    string    `gorm:"type:varchar(255);column:image_url"`
	IsVerified  bool      `gorm:"default:false;column:is_verified"`
	CreatedAt   time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

func EntityToSchema(entity user.User) User {
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
	}
}

func SchemaToEntity(schema User) user.User {
	return user.User{
		ID:          identity.NewIDFromSchema(schema.ID),
		Name:        schema.Name,
		Email:       schema.Email,
		PhoneNumber: schema.PhoneNumber,
		Password:    user.NewPasswordFromSchema(schema.Password),
		Role:        user.NewRoleFromSchema(schema.Role),
		ImageUrl:    shared.NewURLFromSchema(schema.ImageUrl),
		IsVerified:  schema.IsVerified,
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
		},
	}
}
