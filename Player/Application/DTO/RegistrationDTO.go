package DTO

type RegistrationDTO struct {
	Name     string `json:"name" validation:"required"`
	Password string `json:"password" validation:"required"`
}
