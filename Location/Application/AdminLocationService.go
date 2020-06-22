package Application

import (
	"BablarAPI/Geo"
	"BablarAPI/Location/Application/DTO"
	"BablarAPI/Location/Application/Error"
	"BablarAPI/Location/Domain"
	cv "github.com/go-playground/validator"
	"log"
)

type AdminLocationService struct {
	repository Domain.LocationRepository
	log        *log.Logger
}

func NewAdminLocationService(repository Domain.LocationRepository, log *log.Logger) *AdminLocationService {
	return &AdminLocationService{repository: repository, log: log}
}

func (c AdminLocationService) Get(scoordinate string) (error, DTO.MetadataDTO) {
	err, coordinate := Geo.ParseString(scoordinate)
	if err != nil {
		return Error.NewDataValidationError(err), DTO.MetadataDTO{}
	}

	err, location := c.repository.Get(coordinate)

	if err != nil {
		c.log.Print(err)
	}
	return err, DTO.MetadataDTO{
		Name: location.Name,
		Type: location.Type,
	}
}

func (c AdminLocationService) Add(scoordinate string, dto DTO.CreateDTO) error {
	err, coordinate := Geo.ParseString(scoordinate)
	if err != nil {
		return Error.NewDataValidationError(err)
	}
	if err := cv.New().Struct(dto); err != nil {
		return Error.NewDataValidationError(err)
	}
	err = c.repository.Add(Domain.Location{
		Coordinate: coordinate,
		Name:       dto.Name,
		Type:       dto.Type,
	})
	if err != nil {
		c.log.Print(err)
	}
	return err
}

func (c AdminLocationService) Update(scoordinate string, dto DTO.UpdateDTO) error {
	err, coordinate := Geo.ParseString(scoordinate)
	if err != nil {
		return Error.NewDataValidationError(err)
	}
	if err := cv.New().Struct(dto); err != nil {
		return Error.NewDataValidationError(err)
	}
	command:= Domain.UpdateLocationCommand{
		Coordinate:   coordinate,
		Name:         &dto.Name,
		LocationType: &dto.Type,
	}
	if !dto.UpdateName {
		command.Name =nil
	}
	if !dto.UpdateType {
		command.LocationType = nil
	}
	err = c.repository.Update(command)
	if err != nil {
		c.log.Print(err)
	}
	return err
}

func (c AdminLocationService) Delete(scoordinate string) error {
	err, coordinate := Geo.ParseString(scoordinate)
	if err != nil {
		return Error.NewDataValidationError(err)
	}
	err = c.repository.Delete(coordinate)
	if err != nil {
		c.log.Print(err)
	}
	return err
}
