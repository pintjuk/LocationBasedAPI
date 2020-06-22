package Domain

import (
	"github.com/google/uuid"
)

type UpdatePlayerCommand struct {
	Id         uuid.UUID
	Location   *Location
	CashedName *string
}
