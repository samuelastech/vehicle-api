package domain

type Dimensions struct {
	Width, Height, Length float64
}

type Vehicle struct {
	ID               int
	Brand            string
	Model            string
	Color            string
	FabricationYear  int
	MaxSpeed         float64
	SeatingCapacity  int
	FuelType         string
	TransmissionType string
	Registration     string
	Weight           float64
	Dimensions
}

type Vehicles []Vehicle
