package DTO

type CreateDTO struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}
