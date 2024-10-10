package vehicle

import "github.com/samuelastech/vehicle-api/internal/domain"

type repository struct {
	data   domain.Vehicles
	lastId int
}

func NewRepository(data domain.Vehicles) *repository {
	return &repository{data, len(data)}
}

func (r *repository) GetAll() domain.Vehicles {
	return r.data
}

func (r *repository) Create(vehicle domain.Vehicle) (createdEntity domain.Vehicle) {
	r.lastId++
	vehicle.ID = r.lastId
	r.data = append(r.data, vehicle)
	createdEntity = vehicle
	return
}
