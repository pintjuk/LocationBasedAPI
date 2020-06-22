package Domain

import (
	"BablarAPI/Geo"
)

type Location struct {
	Coordinate Geo.Coordinate
	Name       string
	Type       string
}
