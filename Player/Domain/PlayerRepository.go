package Domain

import (
	"github.com/google/uuid"
)

type PlayerRepository interface {
	Get(coordinate uuid.UUID) (error, Player)
	Update(command UpdatePlayerCommand) error
	Add(client Player) error
}
