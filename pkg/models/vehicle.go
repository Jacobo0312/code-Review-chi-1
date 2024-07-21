package models

import (
	"errors"
)

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64
	// Length is the length of the dimension
	Length float64
	// Width is the width of the dimension
	Width float64
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleAttributes struct {
	// Brand is the brand of the vehicle
	Brand string
	// Model is the model of the vehicle
	Model string
	// Registration is the registration of the vehicle
	Registration string
	// Color is the color of the vehicle
	Color string
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int
	// Capacity is the capacity of people of the vehicle
	Capacity int
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64
	// FuelType is the fuel type of the vehicle
	FuelType string
	// Transmission is the transmission of the vehicle
	Transmission string
	// Weight is the weight of the vehicle
	Weight float64
	// Dimensions is the dimensions of the vehicle
	Dimensions
}

// Vehicle is a struct that represents a vehicle in JSON format
type VehicleDoc struct {
	ID              int     `json:"id"`
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

// Vehicle is a struct that represents a vehicle
type Vehicle struct {
	// Id is the unique identifier of the vehicle
	Id int

	// VehicleAttribute is the attributes of a vehicle
	VehicleAttributes
}

// Validate checks if the vehicle data is valid
func (v *Vehicle) Validate() error {
	// Check mandatory fields
	if v.Brand == "" || v.Model == "" || v.Registration == "" {
		return errors.New("los campos Brand, Model y Registration son obligatorios")
	}

	// Additional custom validations if needed
	if v.FabricationYear < 1900 || v.FabricationYear > 2100 {
		return errors.New("año de fabricación fuera de rango válido")
	}

	// Example of a specific validation using dimensions
	if v.Dimensions.Height <= 0 || v.Dimensions.Length <= 0 || v.Dimensions.Width <= 0 {
		return errors.New("dimensiones inválidas")
	}

	if v.Capacity == 0 {
		return errors.New("el atributo pasajers no puede ser 0")
	}

	return nil
}
