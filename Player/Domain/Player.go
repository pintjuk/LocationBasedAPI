package Domain

import (
	"github.com/google/uuid"
)

type Player struct {
	Id           uuid.UUID
	CashedName   string
	LastLocation Location
	Score        int
}
