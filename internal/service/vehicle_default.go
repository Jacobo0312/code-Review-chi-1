package service

import (
	"app/internal/repository"
	"app/pkg/models"
	"errors"
	"strconv"
)

var (
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
	ErrVehiclesNotFound     = errors.New("no se encontraron vehículos con esos criterios.")
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp repository.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp repository.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that add vehicle to map returns a error if vehicle already exist
func (s *VehicleDefault) Create(v models.VehicleDoc) (err error) {
	newVehicle := models.Vehicle{
		Id: v.ID,
		VehicleAttributes: models.VehicleAttributes{
			Brand:           v.Brand,
			Model:           v.Model,
			Registration:    v.Registration,
			Color:           v.Color,
			FabricationYear: v.FabricationYear,
			Capacity:        v.Capacity,
			MaxSpeed:        v.MaxSpeed,
			FuelType:        v.FuelType,
			Transmission:    v.Transmission,
			Weight:          v.Weight,
			Dimensions: models.Dimensions{
				Height: v.Height,
				Length: v.Length,
				Width:  v.Width,
			},
		},
	}

	if err := newVehicle.Validate(); err != nil {
		return err
	}

	err = s.rp.Create(newVehicle)
	if err != nil {
		return ErrVehicleAlreadyExists
	}

	return nil
}

func (s *VehicleDefault) GetByBrandAndYear(brand string, startYear string, endYear string) (v map[int]models.Vehicle, err error) {

	start, err := strconv.Atoi(startYear)

	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(endYear)

	if err != nil {
		return nil, err
	}

	v = s.rp.GetByBrandAndYear(brand, start, end)

	if len(v) == 0 {
		return nil, ErrVehiclesNotFound
	}

	return v, nil

}
