package httpjson

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hferr/device-manager/internal/api/device"
	e "github.com/hferr/device-manager/internal/api/err"
	"github.com/hferr/device-manager/utils/validator"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary      List all devices
// @Description  Get a list of all devices in the system
// @Tags         devices
// @Produce      json
// @Success      200  {array}   device.DTO
// @Failure      500  {object}  err.Error
// @Router       /devices [get]
func (h Handler) ListDevices(w http.ResponseWriter, r *http.Request) {
	ds, err := h.deviceSvs.ListDevices()
	if err != nil {
		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	if err := json.NewEncoder(w).Encode(ds.ToDto()); err != nil {
		e.ServerError(w, e.JSONEncodeErrResp)
		return
	}
}

// @Summary      Create a new device
// @Description  Create a new device in the system.
// @Tags         devices
// @Accept       json
// @Produce      json
// @Param        device  body      device.CreateDeviceRequest  true  "Create device request object"
// @Success      201     {object}  device.DTO
// @Failure      400     {object}  err.Error
// @Failure      422     {object}  err.Errors
// @Failure      500     {object}  err.Error
// @Router       /devices [post]
func (h Handler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	input := device.CreateDeviceRequest{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		e.BadRequest(w, e.JSONDecodeErrResp)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		res, err := json.Marshal(validator.ErrResponse(err))
		if err != nil {
			e.BadRequest(w, e.JSONDecodeErrResp)
			return
		}

		e.UnprocessableEntity(w, res)
		return
	}

	d, err := h.deviceSvs.CreateDevice(input)
	if err != nil {
		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(d.ToDto()); err != nil {
		e.ServerError(w, e.JSONEncodeErrResp)
		return
	}
}

// @Summary      Update a device
// @Description  Update an existing device by its ID, only devices that are not in the
// @Description  state 'in_use' can be updated.
// @Tags         devices
// @Accept       json
// @Produce      json
// @Param        id      path      string                    true  "Device ID"
// @Param        device  body      device.UpdateDeviceRequest  true  "Updated device request object"
// @Success      204
// @Failure      400     {object}  err.Error
// @Failure		 404     {object}  err.Error
// @Failure      422     {object}  err.Errors
// @Failure      500     {object}  err.Error
// @Router       /devices/{id} [patch]
func (h Handler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.InvalidIDErrResp)
		return
	}

	input := device.UpdateDeviceRequest{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		e.BadRequest(w, e.JSONDecodeErrResp)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		res, err := json.Marshal(validator.ErrResponse(err))
		if err != nil {
			e.BadRequest(w, e.JSONDecodeErrResp)
			return
		}

		e.UnprocessableEntity(w, res)
		return
	}

	if err := h.deviceSvs.UpdateDevice(ID, input); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.NotFound(w, e.DeviceNotFoundErrResp)
			return
		}
		if errors.Is(err, device.ErrDeviceInUse) {
			e.UnprocessableEntity(w, e.DeviceInUseErrResp)
			return
		}

		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary      Get device by ID
// @Description  Get a single device by its ID
// @Tags         devices
// @Produce      json
// @Param        id   path      string  true  "Device ID"
// @Success      200  {object}  device.DTO
// @Failure      400  {object}  err.Error
// @Failure      404  {object}  err.Error
// @Failure      500  {object}  err.Error
// @Router       /devices/{id} [get]
func (h Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.InvalidIDErrResp)
		return
	}

	d, err := h.deviceSvs.FindByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.NotFound(w, e.DeviceNotFoundErrResp)
			return
		}

		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	if err := json.NewEncoder(w).Encode(d.ToDto()); err != nil {
		e.ServerError(w, e.JSONEncodeErrResp)
		return
	}
}

// @Summary      Find devices by state
// @Description  Get all devices with a specific state
// @Tags         devices
// @Produce      json
// @Param        state  path      string  true  "Device state"
// @Success      200  {object}  device.DTO
// @Failure      400  {object}  err.Error
// @Failure      404  {object}  err.Error
// @Failure      500  {object}  err.Error
// @Router       /devices/state/{state} [get]
func (h Handler) FindByState(w http.ResponseWriter, r *http.Request) {
	state := chi.URLParam(r, "state")

	ds, err := h.deviceSvs.FindByState(state)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.NotFound(w, e.DeviceNotFoundErrResp)
			return
		}

		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	if err := json.NewEncoder(w).Encode(ds.ToDto()); err != nil {
		e.ServerError(w, e.JSONEncodeErrResp)
		return
	}
}

// @Summary      Find devices by brand
// @Description  Get all devices from a specific brand
// @Tags         devices
// @Produce      json
// @Param        brand  path      string  true  "Device brand"
// @Success      200  {object}  device.DTO
// @Failure      400  {object}  err.Error
// @Failure      404  {object}  err.Error
// @Failure      500  {object}  err.Error
// @Router       /devices/brand/{brand} [get]
func (h Handler) FindByBrand(w http.ResponseWriter, r *http.Request) {
	state := chi.URLParam(r, "brand")

	ds, err := h.deviceSvs.FindByBrand(state)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.NotFound(w, e.DeviceNotFoundErrResp)
			return
		}

		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	if err := json.NewEncoder(w).Encode(ds.ToDto()); err != nil {
		e.ServerError(w, e.JSONEncodeErrResp)
		return
	}
}

// @Summary      Delete a device
// @Description  Delete a device by its ID
// @Tags         devices
// @Produce      json
// @Param        id   path      string  true  "Device ID"
// @Success      204
// @Failure      400  {object}  err.Error
// @Failure		 404  {object}  err.Error
// @Failure      422  {object}  err.Error
// @Failure      500  {object}  err.Error
// @Router       /devices/{id} [delete]
func (h Handler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.InvalidIDErrResp)
		return
	}

	if err := h.deviceSvs.DeleteDevice(ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.NotFound(w, e.DeviceNotFoundErrResp)
			return
		}
		if errors.Is(err, device.ErrDeviceInUse) {
			e.UnprocessableEntity(w, e.DeviceInUseErrResp)
			return
		}

		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
