package Player

import (
	idpApp "BablarAPI/IdP/Application"
	"BablarAPI/IdP/Data/Errors"
	"BablarAPI/Location/Application/Error"
	"BablarAPI/Player/Application"
	"BablarAPI/Player/Application/DTO"
	playerDataErrors "BablarAPI/Player/Data/Errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

func errorsToHTTPStatus(c echo.Context, err error) error {
	switch err.(type) {
	case Errors.NotFound:
		return c.String(http.StatusNotFound, "Not found")
	case playerDataErrors.NotFound:
		return c.String(http.StatusNotFound, "Not found")
	case Error.DataValidationError:
		return c.String(http.StatusBadRequest, "Invalid data")
	case idpApp.UnAuthorizedError:
		return c.String(http.StatusUnauthorized, "UnAuthorized")
	default:
		return c.String(http.StatusInternalServerError, "Internal server error")
	}
}
func InitPlayerClient(g *echo.Group, playerClientService *Application.PlayerClientService) {
	g.PUT("/location", func(c echo.Context) error {
		JWTtoken := c.Get("user").(*jwt.Token)
		location := DTO.LocationDTO{}
		if err := c.Bind(&location); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		err := playerClientService.SendLocation(JWTtoken, location)
		if err == nil {
			return c.JSON(http.StatusOK, nil)
		}

		return errorsToHTTPStatus(c, err)
	})
	g.PUT("/Name", func(c echo.Context) error {
		JWTtoken := c.Get("user").(*jwt.Token)
		body := new(Application.UpdateNameDTO)
		if err := c.Bind(body); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}

		err := playerClientService.UpdateName(JWTtoken, *body)

		if err == nil {
			return c.JSON(http.StatusOK, nil)
		}
		return errorsToHTTPStatus(c, err)
	})

	g.GET("", func(c echo.Context) error {
		JWTtoken := c.Get("user").(*jwt.Token)
		err, player := playerClientService.Get(JWTtoken)
		if err == nil {
			return c.JSON(http.StatusOK, player)
		}
		return errorsToHTTPStatus(c, err)
	})

}

func InitRegistration(g *echo.Group, playerClientService *Application.PlayerClientService) {
	g.POST("", func(c echo.Context) error {
		body := DTO.RegistrationDTO{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		err := playerClientService.New(body)
		if err == nil {
			return c.String(http.StatusOK, "Created!")
		}
		return errorsToHTTPStatus(c, err)
	})
}
func InitLogon(e *echo.Group, idpService *idpApp.IdPService) {
	e.POST("", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		err, token := idpService.Login(username, password)
		if err == nil {
			return c.JSON(http.StatusOK, map[string]string{
				"token": token,
			})
		}
		return c.JSON(http.StatusUnauthorized, nil)
	})
}
