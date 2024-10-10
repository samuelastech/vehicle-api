package vehicle

import (
	"errors"
	"net/http"

	"github.com/samuelastech/vehicle-api/pkg/common"
	"github.com/samuelastech/vehicle-api/pkg/web"
)

type handler struct {
	service common.VechicleService
}

func NewHandler(service common.VechicleService) *handler {
	return &handler{service}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	vehicles, err := h.service.GetAll()
	message := "all vehicles listed"
	code := http.StatusOK

	if err != nil {
		switch {
		case errors.Is(err, ErrVehiclesNotFound):
			code = http.StatusNotFound
			message = err.Error()
		default:
			code = http.StatusInternalServerError
			message = "internal server error"
		}
	}

	web.ResponseJSON(w, code, common.Params{
		"message": message,
		"data":    vehicles,
	})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

}
