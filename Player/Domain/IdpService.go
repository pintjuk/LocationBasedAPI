package Domain

import "github.com/google/uuid"

type IdPService interface {
	GetName(id uuid.UUID) (error, string)
	SetName(id uuid.UUID, name string) error
	New(username string, password string) (error, uuid.UUID)
}
