package httpjson

import (
	"github/hferr/device-manager/internal/api/device"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	deviceSvs device.DeviceService
	validator *validator.Validate
}

func NewHandler(deviceSvs device.DeviceService, v *validator.Validate) *Handler {
	return &Handler{
		deviceSvs: deviceSvs,
		validator: v,
	}
}

func (h Handler) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/devices", func(r chi.Router) {
		r.Get("/", h.ListDevices)
		r.Get("/{id}", h.FindByID)
		r.Post("/", h.CreateDevice)
		r.Patch("/{id}", h.UpdateDevice)
		r.Delete("/{id}", h.DeleteDevice)

		r.Get("/state/{state}", h.FindByState)
		r.Get("/brand/{brand}", h.FindByBrand)
	})

	return r
}
