package domain

type Dimensions struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Length float64 `json:"length"`
}

type Vehicle struct {
	ID               int     `json:"id"`
	Brand            string  `json:"brand"`
	Model            string  `json:"model"`
	Color            string  `json:"color"`
	FabricationYear  int     `json:"fabricationYear"`
	MaxSpeed         float64 `json:"maxSpeed"`
	SeatingCapacity  int     `json:"seatingCapacity"`
	FuelType         string  `json:"fuelType"`
	TransmissionType string  `json:"transmissionType"`
	Registration     string  `json:"registration"`
	Weight           float64 `json:"weight"`
	Dimensions
}

type Vehicles = []Vehicle
