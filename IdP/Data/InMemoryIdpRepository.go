package Data

import (
	"BablarAPI/IdP/Data/Errors"
	"BablarAPI/IdP/Domain"
	"github.com/google/uuid"
)

type inMemoryIdpRepository struct {
	data map[uuid.UUID]Domain.Principal
}

func NewEmptyInMemoryIdpRepository() *inMemoryIdpRepository {
	return &inMemoryIdpRepository{data: make(map[uuid.UUID]Domain.Principal)}
}

func (i inMemoryIdpRepository) GetByUsername(username string) (error, Domain.Principal) {
	for _, value := range i.data {
		if value.Username == username {
			return nil, value
		}
	}
	return Errors.NotFound{Reccord: username}, Domain.Principal{}
}

func (i inMemoryIdpRepository) UpdateUsername(id uuid.UUID, name string) error {
	identity, ok := i.data[id]
	if !ok {
		return Errors.NotFound{Reccord: id.String()}
	}

	identity.Username = name
	i.data[id] = identity
	return nil
}

func (i inMemoryIdpRepository) Get(id uuid.UUID) (error, Domain.Principal) {
	identity, ok := i.data[id]
	if !ok {
		return Errors.NotFound{
			Reccord: id.String(),
		}, Domain.Principal{}
	}
	return nil, identity
}

func (i inMemoryIdpRepository) Add(identity Domain.Principal) error {
	_, ok := i.data[identity.Id]
	if ok {
		return nil
	}
	i.data[identity.Id] = identity
	return nil
}
