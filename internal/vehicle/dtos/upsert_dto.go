package dtos

import "github.com/samuelastech/vehicle-api/internal/domain"

type UpsertDTO struct {
	Brand            string   `json:"brand" validate:"required"`
	Model            string   `json:"model" validate:"required"`
	Color            string   `json:"color" validate:"required"`
	Registration     string   `json:"registration" validate:"required"`
	FuelType         string   `json:"fuelType" validate:"required"`
	TransmissionType string   `json:"transmissionType" validate:"required"`
	FabricationYear  int      `json:"fabricationYear" validate:"required,number,gte=1886"`
	MaxSpeed         float64  `json:"maxSpeed" validate:"required,number,gte=0"`
	SeatingCapacity  int      `json:"seatingCapacity" validate:"required,number,gte=1"`
	Weight           float64  `json:"weight" validate:"required,number,gte=0"`
	Width            float64  `json:"width" validate:"required,number,gte=0"`
	Height           float64  `json:"height" validate:"required,number,gte=0"`
	Length           *float64 `json:"length" validate:"required,number,gte=0"`
}

func (u *UpsertDTO) ToDomain() domain.Vehicle {
	return domain.Vehicle{
		Brand:            u.Brand,
		Model:            u.Model,
		Color:            u.Color,
		Registration:     u.Registration,
		FuelType:         u.FuelType,
		TransmissionType: u.TransmissionType,
		FabricationYear:  u.FabricationYear,
		MaxSpeed:         u.MaxSpeed,
		SeatingCapacity:  u.SeatingCapacity,
		Dimensions: domain.Dimensions{
			Width:  u.Width,
			Height: u.Height,
			Length: *u.Length,
		},
	}
}
