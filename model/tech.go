package model

import "github.com/go-playground/validator/v10"

type Category string

const (
	Frontend    Category = "Frontend"
	Apps        Category = "Apps"
	Backend     Category = "Backend"
	Devops      Category = "Dev/Ops"
	Application Category = "Application"
	Consulting  Category = "Consulting"
	IT_Service  Category = "IT_Service"
)

type Status string

const (
	Core   Status = "core"
	Assess Status = "assess"
	Trial  Status = "trial"
	Adopt  Status = "adopt"
	Hold   Status = "hold"
)

type Tech struct {
	Id          uint     `json:"id" gorm:"primary_key"`
	Category    Category `json:"category" binding:"required" validate:"oneof=Frontend Apps Backend Dev/Ops Consulting Application IT_Service"`
	Status      Status   `json:"status" binding:"required" validate:"oneof=core assess hold adopt trial"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Active      bool     `json:"active"`
	Moved       int8     `json:"moved"`
}

func (t *Tech) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
