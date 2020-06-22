package DTO

type LocationDTO struct {
	Longitude float32 `json:"long" validate:"required"`
	Latitude  float32 `json:"lat" validate:"required"`
}
