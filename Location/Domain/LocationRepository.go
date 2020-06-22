package Domain

import "BablarAPI/Geo"

type LocationRepository interface {
	Get(coordinate Geo.Coordinate) (error, Location)
	Update(coordinate UpdateLocationCommand) error
	Add(Location) error
	Delete(coordinate Geo.Coordinate) error
}
