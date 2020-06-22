package Application

import (
	"BablarAPI/Player/Application/DTO"
	"BablarAPI/Player/Application/Errors"
	"BablarAPI/Player/Domain"
	"errors"
	"github.com/dgrijalva/jwt-go"
	cv "github.com/go-playground/validator"
	"github.com/google/uuid"
	"log"
)

func parseToken(token *jwt.Token) (uuid.UUID, error) {
	sid, ok := token.Claims.(jwt.MapClaims)["id"].(string)
	if !ok {
		return uuid.UUID{}, Errors.NewDataValidationError(errors.New("claims did not contain id"))
	}
	id, err := uuid.Parse(sid)
	if err != nil {
		return uuid.UUID{}, Errors.NewDataValidationError(err)
	}
	return id, nil
}

type PlayerClientService struct {
	playerDomainService *Domain.PlayerService
	errorLogger         *log.Logger
}

func (s PlayerClientService) Get(token *jwt.Token) (error, DTO.PlayerDTO) {
	id, err := parseToken(token)
	if err != nil {
		s.errorLogger.Print(err)
		return err, DTO.PlayerDTO{}
	}
	err, player := s.playerDomainService.GetPlayer(id)
	if err != nil {
		s.errorLogger.Print(err)
	}
	return err, DTO.MakeDTOFromPlayer(player)
}

type UpdateNameDTO struct {
	Name string `json:"name" validate:"required"`
}

func (s PlayerClientService) UpdateName(token *jwt.Token, dto UpdateNameDTO) error {
	if err:= cv.New().Struct(dto); err != nil{
		return Errors.NewDataValidationError(err)
	}
	id, err := parseToken(token)
	if err != nil {
		s.errorLogger.Print(err)
		return err
	}
	err = s.playerDomainService.UpdateName(id, dto.Name)
	if err != nil {
		s.errorLogger.Print(err)
	}
	return err
}

func (s PlayerClientService) SendLocation(token *jwt.Token, dto DTO.LocationDTO) error {
	if err := cv.New().Struct(dto); err != nil {
		err = Errors.NewDataValidationError(err)
		s.errorLogger.Print(err)
		return err
	}
	id, err := parseToken(token)
	if err != nil {
		err = Errors.NewDataValidationError(err)
		s.errorLogger.Print(err)
		return err
	}
	err = s.playerDomainService.SendLocation(id, Domain.Location{dto.Longitude, dto.Latitude})
	if err != nil {
		s.errorLogger.Print(err)
	}
	return err
}

func (s PlayerClientService) New(dto DTO.RegistrationDTO) error {
	if err := cv.New().Struct(dto); err != nil {
		err = Errors.NewDataValidationError(err)
		s.errorLogger.Print(err)
		return Errors.DataValidationError{}
	}
	err := s.playerDomainService.New(dto.Name, dto.Password)
	if err != nil {
		s.errorLogger.Print(err)
	}
	return err
}

func NewPlayerClientService(playerService *Domain.PlayerService) *PlayerClientService {
	return &PlayerClientService{playerDomainService: playerService}
}
