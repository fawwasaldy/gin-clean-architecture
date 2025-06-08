package refresh_token

import (
	"github.com/google/uuid"
	"kpl-base/domain/identity"
	"kpl-base/domain/refresh_token"
	"kpl-base/domain/shared"
	"time"
)

type RefreshToken struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();column:id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;column:user_id"`
	Token     string    `gorm:"type:varchar(255);not null;uniqueIndex;column:token"`
	ExpiresAt time.Time `gorm:"type:timestamp with time zone;not null;column:expires_at"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone;column:updated_at"`
}

func EntityToSchema(entity refresh_token.RefreshToken) RefreshToken {
	return RefreshToken{
		ID:        entity.ID.ID,
		UserID:    entity.UserID.ID,
		Token:     entity.Token,
		ExpiresAt: entity.ExpiresAt,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
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
		},
	}
}
