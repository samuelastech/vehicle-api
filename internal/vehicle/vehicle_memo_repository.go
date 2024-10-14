package vehicle

import (
	"strings"

	"github.com/samuelastech/vehicle-api/internal/domain"
	"github.com/samuelastech/vehicle-api/pkg/common"
)

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

func (r *repository) GetAllBy(params common.Params, modifier int) (entities domain.Vehicles) {
	for _, vehicle := range r.data {
		var paramsMatched int

		for key, value := range params {
			switch key {
			case "yearEqual":
				if vehicle.FabricationYear == value {
					if modifier == Or {
						entities = append(entities, vehicle)
						break
					} else {
						paramsMatched++
					}
				}

			case "colorEqual":
				if color, ok := value.(string); ok {
					if strings.ToLower(vehicle.Color) == strings.ToLower(color) {
						if modifier == Or {
							entities = append(entities, vehicle)
							break
						} else {
							paramsMatched++
						}
					}
				}
			}
		}

		if paramsMatched == len(params) {
			entities = append(entities, vehicle)
		}
	}

	return
}

func (r *repository) Create(vehicle domain.Vehicle) (createdEntity domain.Vehicle) {
	r.lastId++
	vehicle.ID = r.lastId
	r.data = append(r.data, vehicle)
	createdEntity = vehicle
	return
}
