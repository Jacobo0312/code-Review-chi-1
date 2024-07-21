package repository

import (
	"app/pkg/models"
	"errors"
)

var (
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]models.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]models.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]models.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that returns a Vehicle struct
func (r *VehicleMap) Create(v models.Vehicle) (err error) {
	// Get id
	id := v.Id

	//Verify if id already exist
	_, ok := r.db[id]

	if ok {
		return ErrVehicleAlreadyExists
	}

	//Insert
	r.db[id] = v

	return nil
}
