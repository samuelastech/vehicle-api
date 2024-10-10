package common

import "github.com/samuelastech/vehicle-api/internal/domain"

type Repository interface {
	GetAll() domain.Vehicles
}
