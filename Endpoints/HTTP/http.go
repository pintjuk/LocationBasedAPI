package HTTP

import (
	"BablarAPI/Config"
	"BablarAPI/Endpoints/HTTP/Location"
	PlayerEndpoints "BablarAPI/Endpoints/HTTP/Player"
	IdP "BablarAPI/IdP/Application"
	LocationApp "BablarAPI/Location/Application"
	PlayerApp "BablarAPI/Player/Application"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHttp(IDPService *IdP.IdPService, PlayerService *PlayerApp.PlayerClientService, CLS *LocationApp.ClientLocationService, ALS *LocationApp.AdminLocationService) {
	// TODO: create Configuration managment for paths
	// TODO: Decide weather to return some message  upon errors,
	//       right now we return null as mesage, however with correct http status codes
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Requier Authenticaion for client group
	clientGroup := e.Group("/Client")
	clientGroup.Use(middleware.JWT([]byte(Config.SECRET)))
	PlayerEndpoints.InitPlayerClient(clientGroup.Group("/Player"), PlayerService)
	Location.InitLocationClient(clientGroup.Group("/Location"), CLS)

	// End points that dont requir outhentication
	PlayerEndpoints.InitLogon(e.Group("/Client/Login"), IDPService)
	PlayerEndpoints.InitRegistration(e.Group("/Client/Register"), PlayerService)
	adminGroup := e.Group("/Admin")
	Location.InitLocationAdmin(adminGroup.Group("/Location"), ALS)

	e.Logger.Fatal(e.Start(":1234"))
}
