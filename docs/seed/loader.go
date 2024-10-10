package seed

import (
	"encoding/json"
	"os"

	"github.com/samuelastech/vehicle-api/internal/domain"
)

type VehicleJSON struct {
	Id              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

func VehicleJSONFile() domain.Vehicles {
	file, err := os.Open("docs/seed/vehicles_100.json")

	if err != nil {
		return nil
	}

	defer file.Close()
	var vehiclesJSON []VehicleJSON
	err = json.NewDecoder(file).Decode(&vehiclesJSON)
	if err != nil {
		return nil
	}

	// serialize vehicles
	v := make(domain.Vehicles, 100)

	for i, vh := range vehiclesJSON {
		v[i] = domain.Vehicle{
			ID:               vh.Id,
			Brand:            vh.Brand,
			Model:            vh.Model,
			Registration:     vh.Registration,
			Color:            vh.Color,
			FabricationYear:  vh.FabricationYear,
			SeatingCapacity:  vh.Capacity,
			MaxSpeed:         vh.MaxSpeed,
			FuelType:         vh.FuelType,
			TransmissionType: vh.Transmission,
			Weight:           vh.Weight,
			Dimensions: domain.Dimensions{
				Height: vh.Height,
				Length: vh.Length,
				Width:  vh.Width,
			},
		}
	}

	return v
}
