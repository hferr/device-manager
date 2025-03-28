package httpjson

import (
	"github/hferr/device-manager/internal/api/device"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	deviceSvs device.DeviceService
}

func NewHandler(deviceSvs device.DeviceService) *Handler {
	return &Handler{
		deviceSvs: deviceSvs,
	}
}

func (h Handler) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return r
}
