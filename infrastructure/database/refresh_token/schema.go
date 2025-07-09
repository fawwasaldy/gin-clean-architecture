package refresh_token

import (
	"gin-clean-architecture/domain/identity"
	"gin-clean-architecture/domain/refresh_token"
	"gin-clean-architecture/domain/shared"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;column:user_id"`
	Token     string         `gorm:"type:varchar(255);not null;uniqueIndex;column:token"`
	ExpiresAt time.Time      `gorm:"type:timestamp with time zone;not null;column:expires_at"`
	CreatedAt time.Time      `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp with time zone;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp with time zone;column:deleted_at"`
}

func EntityToSchema(entity refresh_token.RefreshToken) RefreshToken {
	var deletedAtTime time.Time
	if entity.Timestamp.DeletedAt != nil {
		deletedAtTime = *entity.Timestamp.DeletedAt
	} else {
		deletedAtTime = time.Time{}
	}
	return RefreshToken{
		ID:        entity.ID.ID,
		UserID:    entity.UserID.ID,
		Token:     entity.Token,
		ExpiresAt: entity.ExpiresAt,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: gorm.DeletedAt{
			Time:  deletedAtTime,
			Valid: entity.Timestamp.DeletedAt != nil,
		},
	}
}

func SchemaToEntity(schema RefreshToken) refresh_token.RefreshToken {
	return refresh_token.RefreshToken{
		ID:        identity.NewIDFromSchema(schema.ID),
		UserID:    identity.NewIDFromSchema(schema.UserID),
		Token:     schema.Token,
		ExpiresAt: schema.ExpiresAt,
		Timestamp: shared.Timestamp{
			CreatedAt: schema.CreatedAt,
			UpdatedAt: schema.UpdatedAt,
			DeletedAt: &schema.DeletedAt.Time,
		},
	}
}
