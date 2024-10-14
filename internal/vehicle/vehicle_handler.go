package vehicle

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
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

func (h *handler) GetByColorAndYear(w http.ResponseWriter, r *http.Request) {
	color := chi.URLParam(r, "color")

	if !common.IsWord(color) {
		web.ResponseJSON(w, http.StatusBadRequest, common.Params{
			"message": "'color' param is not valid",
		})
		return
	}

	color = strings.ToLower(color)
	year, err := strconv.Atoi(chi.URLParam(r, "year"))

	if err != nil {
		web.ResponseJSON(w, http.StatusBadRequest, common.Params{
			"message": "'year' param is not valid",
		})
		return
	}

	vehicles, err := h.service.GetAllBy(common.Params{"yearEqual": year, "colorEqual": color}, And)

	if err != nil {
		code := http.StatusInternalServerError
		message := "internal server error"

		switch {
		case errors.Is(err, ErrVehiclesMatchingParamsNotFound):
			code = http.StatusNotFound
			message = "there are not any vehicles matching those params"
		}

		web.ResponseJSON(w, code, common.Params{
			"message": message,
			"data":    nil,
		})

		return
	}

	web.ResponseJSON(w, http.StatusOK, common.Params{
		"message": "showing vehicles with specific year and color",
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
