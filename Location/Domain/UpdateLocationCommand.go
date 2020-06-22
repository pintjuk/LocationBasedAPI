package Domain

import "BablarAPI/Geo"

type UpdateLocationCommand struct {
	Coordinate   Geo.Coordinate
	Name         *string
	LocationType *string
}
