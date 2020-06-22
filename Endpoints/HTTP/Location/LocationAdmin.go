package Location

import (
	"BablarAPI/Location/Application"
	"BablarAPI/Location/Application/DTO"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

func InitLocationAdmin(g *echo.Group, locationService *Application.AdminLocationService) {

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

	g.POST("/:coordinate", func(c echo.Context) error {
		coordinate := c.Param("coordinate")
		coordinate, err:= url.QueryUnescape(coordinate)
		if err!=nil{
			return c.String(http.StatusBadRequest, "Bad coordinate")
		}
		dto := DTO.CreateDTO{}
		c.Bind(&dto)
		err = locationService.Add(coordinate, dto)
		if err == nil {
			return c.JSON(http.StatusOK, nil)
		}
		return errorsToHTTPStatus(c, err)
	})

	g.PUT("/:coordinate", func(c echo.Context) error {
		coordinate := c.Param("coordinate")
		coordinate, err:= url.QueryUnescape(coordinate)
		if err!=nil{
			return c.String(http.StatusBadRequest, "Bad coordinate")
		}
		dto := DTO.UpdateDTO{}
		if err := c.Bind(&dto); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		err = locationService.Update(coordinate, dto)
		if err == nil {
			return c.JSON(http.StatusOK, nil)
		}
		return errorsToHTTPStatus(c, err)
	})

	g.DELETE("/:coordinate", func(c echo.Context) error {
		coordinate := c.Param("coordinate")
		coordinate, err:= url.QueryUnescape(coordinate)
		if err!=nil{
			return c.String(http.StatusBadRequest, "Bad coordinate")
		}
		err = locationService.Delete(coordinate)
		if err == nil {
			return c.JSON(http.StatusOK, nil)
		}
		return errorsToHTTPStatus(c, err)
	})
}
