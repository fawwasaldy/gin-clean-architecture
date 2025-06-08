package identity

import "github.com/google/uuid"

type ID struct {
	id uuid.UUID
}

func NewID() ID {
	return ID{
		id: uuid.New(),
	}
}

func (i ID) Id() uuid.UUID {
	return i.id
}

func (i ID) String() string {
	return i.id.String()
}
