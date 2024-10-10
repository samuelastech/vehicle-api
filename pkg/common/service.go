package common

import "github.com/samuelastech/vehicle-api/internal/domain"

type VechicleService interface {
	GetAll() (vehicles domain.Vehicles, err error)
	// GetAllBy(params Params) (vehicles domain.Vehicles, err error)
	Create(vehicle domain.Vehicle) (createdVehicle domain.Vehicle)
}
