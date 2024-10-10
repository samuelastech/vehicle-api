package vehicle

import (
	"github.com/samuelastech/vehicle-api/internal/domain"
	"github.com/samuelastech/vehicle-api/pkg/common"
)

type service struct {
	repository common.Repository[domain.Vehicle]
}

func NewService(repository common.Repository[domain.Vehicle]) *service {
	return &service{repository}
}

func (s *service) GetAll() (domain.Vehicles, error) {
	list := s.repository.GetAll()

	if len(list) <= 0 {
		return nil, ErrVehiclesNotFound
	}

	return list, nil
}

func (s *service) Create(vehicle domain.Vehicle) (createdVehicle domain.Vehicle) {
	createdVehicle = s.repository.Create(vehicle)
	return
}
