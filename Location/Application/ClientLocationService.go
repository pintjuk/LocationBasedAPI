package Application

import (
	"BablarAPI/Geo"
	"BablarAPI/Location/Application/Error"
	"BablarAPI/Location/Domain"
	"log"
)

type ClientLocationService struct {
	repository Domain.LocationRepository
	log        *log.Logger
}

func NewClientLocationService(repository Domain.LocationRepository, log *log.Logger) *ClientLocationService {
	return &ClientLocationService{repository: repository, log: log}
}

func (c ClientLocationService) Get(scoordinate string) (error, Domain.Location) {
	err, coordinate := Geo.ParseString(scoordinate)
	if err != nil {
		return Error.NewDataValidationError(err), Domain.Location{}
	}
	err, location := c.repository.Get(coordinate)
	if err != nil {
		c.log.Print(err)
	}
	return err, location
}
