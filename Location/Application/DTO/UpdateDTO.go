package DTO

type UpdateDTO struct {
	UpdateName bool `json:"updateName"`
	UpdateType bool `json:"updateType"`
	Name string `json:"name" validate:""`
	Type string `json:"type" validate:""`
}
