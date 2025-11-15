package shared

import "time"

type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
