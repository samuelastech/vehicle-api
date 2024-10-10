package vehicle

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/samuelastech/vehicle-api/internal/vehicle/dtos"
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
	var body dtos.UpsertDTO
	json.NewDecoder(r.Body).Decode(&body)
	err := validator.New().Struct(body)
	message := "created the vehicle successfully"
	code := http.StatusCreated

	if err != nil {
		code = http.StatusBadRequest
		message = "there are invalid fields"
		validationErrors, _ := err.(validator.ValidationErrors)
		var errorsList []string

		for _, fail := range validationErrors {
			errorsList = append(errorsList, fmt.Sprintf("field '%s' is not valid", fail.Field()))
		}

		web.ResponseJSON(w, code, common.Params{
			"message": message,
			"errors":  errorsList,
		})

		return
	}

	createdVehicle := h.service.Create(body.ToDomain())
	web.ResponseJSON(w, code, common.Params{
		"message": message,
		"data":    createdVehicle,
	})
}
