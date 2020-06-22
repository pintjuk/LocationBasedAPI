package Domain

import "github.com/google/uuid"

type IdPRepsitory interface {
	GetByUsername(username string) (error, Principal)
	UpdateUsername(id uuid.UUID, name string) error
	Add(Principal) error
	Get(id uuid.UUID) (error, Principal)
}
