package Domain

import "github.com/google/uuid"

type Principal struct {
	Id       uuid.UUID
	Username string
	Password string
}
