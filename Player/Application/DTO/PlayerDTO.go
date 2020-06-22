package DTO

import (
	"BablarAPI/Player/Domain"
	"github.com/google/uuid"
)

type PlayerDTO struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `jsan:"name"`
	Score    int       `json:"score"`
	LastLong float32   `json:"lastLong"`
	LastLat  float32   `json:"lastLat"`
}

func MakeDTOFromPlayer(p Domain.Player) PlayerDTO {
	return PlayerDTO{
		Id:       p.Id,
		Name:     p.CashedName,
		Score:    p.Score,
		LastLong: p.LastLocation.Latitude,
		LastLat:  p.LastLocation.Longitude,
	}
}
