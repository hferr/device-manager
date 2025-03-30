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
		if errors.Is(err, device.ErrDeviceInUse) {
			e.UnprocessableEntity(w, e.DeviceInUseErrResp)
			return
		}

		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

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

func (h Handler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.InvalidIDErrResp)
		return
	}

	if err := h.deviceSvs.DeleteDevice(ID); err != nil {
		if errors.Is(err, device.ErrDeviceInUse) {
			e.UnprocessableEntity(w, e.DeviceInUseErrResp)
			return
		}

		e.ServerError(w, e.DeviceServiceFailedErrResp)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
