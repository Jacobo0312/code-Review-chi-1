package handler

import (
	"app/internal/service"
	appErrors "app/pkg/errors"
	"app/pkg/helpers"
	"app/pkg/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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
			helpers.RespondWithError(w, appErrors.NewBadRequest("Datos del vehículo mal formados o incompletos.", err))
			return
		}

		err := h.sv.Create(newVehicleDoc)
		if err != nil {
			log.Println(err)

			//Comparate  vehicle already exists
			if errors.Is(err, service.ErrVehicleAlreadyExists) {
				helpers.RespondWithError(w, appErrors.NewConflict("Identificador del vehículo ya existente."))
				return
			}

			helpers.RespondWithError(w, appErrors.NewBadRequest("Datos del vehículo mal formados o incompletos.", err))
			return
		}

		helpers.RespondWithJSON(w, http.StatusCreated, newVehicleDoc)

	}
}

//GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
//200 OK: Devuelve una lista de vehículos que cumplen con los criterios.
//404 Not Found: No se encontraron vehículos con esos criterios.

func (h *VehicleDefault) GetByBrandAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		startYear := chi.URLParam(r, "start_year")
		endYear := chi.URLParam(r, "end_year")

		vehicles, err := h.sv.GetByBrandAndYear(brand, startYear, endYear)
		if err != nil {
			helpers.RespondWithError(w, appErrors.NewNotFound("No se encontraron vehículos con esos criterios.", err))
			return
		}

		helpers.RespondWithJSON(w, http.StatusOK, vehicles)

	}
}

// GET /vehicles/color/{color}/year/{year}
// 200 OK: Devuelve una lista de vehículos que cumplen con los criterios.
// 404 Not Found: No se encontraron vehículos con esos criterios.

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")

		v, err := h.sv.GetByColorAndYear(color, year)

		if err != nil {
			if errors.Is(err, service.ErrVehiclesNotFound) {
				helpers.RespondWithError(w, appErrors.NewNotFound("No se encontraron vehículos con esos criterios.", err))
			} else {
				helpers.RespondWithError(w, appErrors.NewBadRequest(err.Error(), err))
			}

			return
		}

		helpers.RespondWithJSON(w, http.StatusAccepted, v)

	}
}
