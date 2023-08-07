package models

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type PostRequest struct {
	Name    string `json:"name" validate:"required,gt=0,lte=200,alpha"`
	Surname string `json:"surname" validate:"required,gt=0,lte=200,alpha"`
	Age     int    `json:"age" validate:"required,gt=0,lte=130,number"`
}

func (p *PostRequest) Validate() error {
	validate = validator.New()
	return validate.Struct(p)
}

type PatchRequest struct {
	ID      string `json:"id" validate:"required"`
	Age     int    `json:"age" validate:"required,gt=0,lte=130,number"`
	Surname string `json:"surname" validate:"required,gt=0,lte=200,alpha"`
}

func (p *PatchRequest) Validate() error {
	validate = validator.New()
	return validate.Struct(p)
}

type DeleteRequest struct {
	ID string `json:"id" validate:"required"`
}

func (d *DeleteRequest) Validate() error {
	validate = validator.New()
	return validate.Struct(d)
}
