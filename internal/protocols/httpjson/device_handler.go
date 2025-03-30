package httpjson

import (
	"encoding/json"
	"github/hferr/device-manager/internal/api/device"
	"github/hferr/device-manager/utils/validator"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (h Handler) ListDevices(w http.ResponseWriter, r *http.Request) {
	ds, err := h.deviceSvs.ListDevices()
	if err != nil {
		http.Error(w, "failed to list devices", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ds.ToDto()); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h Handler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	input := device.CreateDeviceRequest{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		res, err := json.Marshal(validator.ErrResponse(err))
		if err != nil {
			http.Error(w, "failed to marshal error response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	d, err := h.deviceSvs.CreateDevice(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(d.ToDto()); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h Handler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid ID param", http.StatusBadRequest)
		return
	}

	input := device.UpdateDeviceRequest{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(input); err != nil {
		res, err := json.Marshal(validator.ErrResponse(err))
		if err != nil {
			http.Error(w, "failed to marshal error response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	if err := h.deviceSvs.UpdateDevice(ID, input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid ID param", http.StatusBadRequest)
		return
	}

	d, err := h.deviceSvs.FindByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(d.ToDto()); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h Handler) FindByState(w http.ResponseWriter, r *http.Request) {
	state := chi.URLParam(r, "state")

	ds, err := h.deviceSvs.FindByState(state)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ds.ToDto()); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h Handler) FindByBrand(w http.ResponseWriter, r *http.Request) {
	state := chi.URLParam(r, "brand")

	ds, err := h.deviceSvs.FindByBrand(state)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ds.ToDto()); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h Handler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid ID param", http.StatusBadRequest)
		return
	}

	if err := h.deviceSvs.DeleteDevice(ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
