package vehicle

import (
	"github.com/samuelastech/vehicle-api/internal/domain"
	"github.com/samuelastech/vehicle-api/pkg/common"
)

type service struct {
	repository common.Repository
}

func NewService(repository common.Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() (domain.Vehicles, error) {
	list := s.repository.GetAll()

	if len(list) <= 0 {
		return nil, ErrVehiclesNotFound
	}

	return list, nil
}
