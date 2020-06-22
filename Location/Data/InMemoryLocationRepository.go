package Data

import (
	"BablarAPI/Geo"
	"BablarAPI/IdP/Data/Errors"
	"BablarAPI/Location/Domain"
	"errors"
)

type InMemoryLocationRepository struct {
	data map[Geo.Coordinate]Domain.Location
}

func NewInMemoryLocationRepository() *InMemoryLocationRepository {
	return &InMemoryLocationRepository{data: make(map[Geo.Coordinate]Domain.Location)}
}

func (i InMemoryLocationRepository) Get(coordinate Geo.Coordinate) (error, Domain.Location) {
	result, ok := i.data[coordinate]
	if !ok {
		err := errors.New("Record does not exist")
		return err, result
	}
	return nil, result
}

func (i InMemoryLocationRepository) Add(location Domain.Location) error {
	_, ok := i.data[location.Coordinate]
	if ok {
		return nil
	}
	i.data[location.Coordinate] = location
	return nil
}

func (i InMemoryLocationRepository) Delete(coordinate Geo.Coordinate) error {
	delete(i.data, coordinate)
	return nil
}
func (i InMemoryLocationRepository) Update(updateCommand Domain.UpdateLocationCommand) error {
	_, ok := i.data[updateCommand.Coordinate]
	if !ok {
		return Errors.NotFound{Reccord: updateCommand.Coordinate.ToString()}
	}
	if updateCommand.LocationType != nil {
		h := i.data[updateCommand.Coordinate]
		h.Type = *updateCommand.LocationType
		i.data[updateCommand.Coordinate] = h
	}
	if updateCommand.Name != nil {
		h := i.data[updateCommand.Coordinate]
		h.Name = *updateCommand.Name
		i.data[updateCommand.Coordinate] = h
	}
	return nil
}
