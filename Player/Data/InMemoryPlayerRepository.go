package Data

import (
	"BablarAPI/IdP/Data/Errors"
	"BablarAPI/Player/Domain"
	"github.com/google/uuid"
)

type inMemoryPlayerRepository struct {
	data map[uuid.UUID]Domain.Player
}

func NewInMemoryPlayerRepository() *inMemoryPlayerRepository {
	return &inMemoryPlayerRepository{data: make(map[uuid.UUID]Domain.Player)}
}

func (i inMemoryPlayerRepository) Get(clientId uuid.UUID) (error, Domain.Player) {
	result, ok := i.data[clientId]
	if !ok {
		return Errors.NotFound{
			Reccord: clientId.String(),
		}, Domain.Player{}
	}
	return nil, result
}
func (i inMemoryPlayerRepository) Add(client Domain.Player) error {
	_, ok := i.data[client.Id]
	if ok {
		return nil
	}
	i.data[client.Id] = client
	return nil
}
func (i inMemoryPlayerRepository) Update(updateCommand Domain.UpdatePlayerCommand) error {
	_, ok := i.data[updateCommand.Id]
	if !ok {
		return Errors.NotFound{
			Reccord: updateCommand.Id.String(),
		}
	}
	if updateCommand.Location != nil {
		h := i.data[updateCommand.Id]
		h.LastLocation = *updateCommand.Location
		i.data[updateCommand.Id] = h
	}
	if updateCommand.CashedName != nil {
		h := i.data[updateCommand.Id]
		h.CashedName = *updateCommand.CashedName
		i.data[updateCommand.Id] = h
	}
	return nil
}
