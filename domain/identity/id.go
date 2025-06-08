package identity

import "github.com/google/uuid"

type ID struct {
	ID uuid.UUID
}

func NewID(id uuid.UUID) ID {
	return ID{
		ID: id,
	}
}

func NewIDFromSchema(id uuid.UUID) ID {
	return ID{
		ID: id,
	}
}

func (i ID) String() string {
	return i.ID.String()
}
