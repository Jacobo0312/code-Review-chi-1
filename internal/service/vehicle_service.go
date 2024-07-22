package service

import "app/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)

	//Create is a method that add vehicle to map returns a error if vehicle already exist
	Create(v models.VehicleDoc) (err error)

	//GetByBrandAndYear is method that returns a map of vehicles of specific brand and  manufactured over a range of years.
	GetByBrandAndYear(brand string, startYear string, endYear string) (v map[int]models.Vehicle, err error)

	GetByColorAndYear(color string, year string) (v map[int]models.Vehicle, err error)

	GetByRangeWeight(min string, max string) (v map[int]models.Vehicle, err error)
}
