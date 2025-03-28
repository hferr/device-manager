package httpjson

import (
	"encoding/json"
	"github/hferr/device-manager/internal/api/device"
	"github/hferr/device-manager/utils/validator"
	"net/http"
)

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

	if err := json.NewEncoder(w).Encode(d.ToDTO()); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
