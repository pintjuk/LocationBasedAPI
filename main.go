package main

import (
	"BablarAPI/Endpoints/HTTP"
	"BablarAPI/Geo"
	IdP "BablarAPI/IdP/Application"
	IdPData "BablarAPI/IdP/Data"
	IdPDomain "BablarAPI/IdP/Domain"
	"BablarAPI/Location/Application"
	Data2 "BablarAPI/Location/Data"
	Domain2 "BablarAPI/Location/Domain"
	Player "BablarAPI/Player/Application"
	"BablarAPI/Player/Data"
	"BablarAPI/Player/Domain"
	"github.com/google/uuid"
	"log"
	"os"
)

func main() {
	errLogger := log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	IdPRepository := IdPData.NewEmptyInMemoryIdpRepository()
	testUserId := uuid.New()
	_ = IdPRepository.Add(IdPDomain.Principal{
		Id:       testUserId,
		Username: "TestUser1",
		Password: "1234",
	})
	PlayerRepository := Data.NewInMemoryPlayerRepository()
	_ = PlayerRepository.Add(Domain.Player{
		Id:           testUserId,
		CashedName:   "TestUser1",
		LastLocation: Domain.Location{Longitude: 59, Latitude: 18},
		Score:        10,
	})
	IDPService := IdP.NewIdPService(IdPDomain.NewIdPService(IdPRepository), errLogger)
	LocationRepository := Data2.NewInMemoryLocationRepository()
	_ = LocationRepository.Add(Domain2.Location{
		Coordinate: Geo.Coordinate{Longitude: 10, Latitude: 10},
		Name:       "Test location",
		Type:       "Type1",
	})
	HTTP.SetupHttp(
		IDPService,
		Player.NewPlayerClientService(Domain.NewPlayerService(IDPService, PlayerRepository)),
		Application.NewClientLocationService(LocationRepository, errLogger),
		Application.NewAdminLocationService(LocationRepository, errLogger))
}
