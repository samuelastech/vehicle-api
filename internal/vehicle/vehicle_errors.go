package vehicle

import "errors"

var (
	ErrVehiclesNotFound               = errors.New("there are not any vehicles")
	ErrVehiclesMatchingParamsNotFound = errors.New("there are not any vehicles matching these params")
)
