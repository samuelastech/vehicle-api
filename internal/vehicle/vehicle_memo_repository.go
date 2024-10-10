package vehicle

import "github.com/samuelastech/vehicle-api/internal/domain"

type repository struct {
	data domain.Vehicles
}

func NewRepository(data domain.Vehicles) *repository {
	return &repository{data}
}

func (r *repository) GetAll() domain.Vehicles {
	return r.data
}
