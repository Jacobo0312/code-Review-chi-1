package repository

import (
	"app/pkg/models"
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)

	//Create is a method that add vehicle to map returns a error if vehicle already exist
	Create(v models.Vehicle) (err error)

	//GetByBrandAndYear is method that returns a map of vehicles of specific brand and  manufactured over a range of years.
	GetByBrandAndYear(brand string, startYear int, endYear int) (v map[int]models.Vehicle)

	GetByColorAndYear(color string, year int) (v map[int]models.Vehicle)

	GetByRangeWeight(min float64, max float64) (v map[int]models.Vehicle)
}
