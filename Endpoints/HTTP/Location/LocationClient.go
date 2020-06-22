package Location

import (
	"BablarAPI/Location/Application"
	"BablarAPI/Location/Application/Error"
	dataError "BablarAPI/Location/Data/Errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

func errorsToHTTPStatus(c echo.Context, err error) error {
	switch err.(type) {
	case dataError.NotFound:
		return c.JSON(http.StatusNotFound, nil)
	case Error.DataValidationError:
		return c.JSON(http.StatusBadRequest, nil)
	default:
		return c.JSON(http.StatusInternalServerError, nil)
	}
}
func InitLocationClient(g *echo.Group, locationService *Application.ClientLocationService) {
	g.GET("/:coordinate", func(c echo.Context) error {
		coordinate := c.Param("coordinate")
		coordinate, err:= url.QueryUnescape(coordinate)
		if err!=nil{
			return c.String(http.StatusBadRequest, "Bad coordinate")
		}

		err, location := locationService.Get(coordinate)
		if err == nil {
			return c.JSON(http.StatusOK, location)
		}
		return errorsToHTTPStatus(c, err)
	})
}
