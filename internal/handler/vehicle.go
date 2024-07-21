package handler

import (
	"app/internal/service"
	"app/pkg/errors"
	"app/pkg/helpers"
	"app/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv service.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv service.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)
		for key, value := range v {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// 	201 Created: Vehículo creado exitosamente.
// 400 Bad Request: Datos del vehículo mal formados o incompletos.
// 409 Conflict: Identificador del vehículo ya existente.

// Create is a method that add new vehicle, returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newVehicleDoc models.VehicleDoc

		if err := json.NewDecoder(r.Body).Decode(&newVehicleDoc); err != nil {
			helpers.RespondWithError(w, errors.NewBadRequest("Datos del vehículo mal formados o incompletos.", err))
			return
		}

		v := models.Vehicle{
			Id: newVehicleDoc.ID,
			VehicleAttributes: models.VehicleAttributes{
				Brand:           newVehicleDoc.Brand,
				Model:           newVehicleDoc.Model,
				Registration:    newVehicleDoc.Registration,
				Color:           newVehicleDoc.Color,
				FabricationYear: newVehicleDoc.FabricationYear,
				Capacity:        newVehicleDoc.Capacity,
				MaxSpeed:        newVehicleDoc.MaxSpeed,
				FuelType:        newVehicleDoc.FuelType,
				Transmission:    newVehicleDoc.Transmission,
				Weight:          newVehicleDoc.Weight,
				Dimensions: models.Dimensions{
					Height: newVehicleDoc.Height,
					Length: newVehicleDoc.Length,
					Width:  newVehicleDoc.Width,
				},
			},
		}

		if err := v.Validate(); err != nil {
			helpers.RespondWithError(w, errors.NewBadRequest("Datos del vehículo mal formados o incompletos.", err))
			return
		}

		err := h.sv.Create(v)
		if err != nil {
			helpers.RespondWithError(w, errors.NewConflict("Identificador del vehículo ya existente."))
		} else {
			helpers.RespondWithJSON(w, http.StatusCreated, v)

		}

	}
}
