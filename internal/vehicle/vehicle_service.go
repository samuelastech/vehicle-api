package vehicle

import (
	"github.com/samuelastech/vehicle-api/internal/domain"
	"github.com/samuelastech/vehicle-api/pkg/common"
)

type service struct {
	repository common.Repository[domain.Vehicle]
}

const (
	And = 0
	Or  = 1
)

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

func (s *service) GetAllBy(params common.Params, modifier int) (vehicles domain.Vehicles, err error) {
	vehicles = s.repository.GetAllBy(params, modifier)

	if len(vehicles) <= 0 {
		err = ErrVehiclesMatchingParamsNotFound
		return
	}

	return
}
