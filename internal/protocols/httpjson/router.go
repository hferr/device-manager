package httpjson

import (
	"net/http"

	"github.com/hferr/device-manager/internal/api/device"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

const (
	HeaderKeyContentType       = "Content-Type"
	HeaderValueContentTypeJSON = "application/json;charset=utf8"
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
		r.Use(middlewareContentTypeJSON)

		r.Get("/", h.ListDevices)
		r.Get("/{id}", h.FindByID)
		r.Post("/", h.CreateDevice)
		r.Patch("/{id}", h.UpdateDevice)
		r.Delete("/{id}", h.DeleteDevice)

		r.Get("/state/{state}", h.FindByState)
		r.Get("/brand/{brand}", h.FindByBrand)
	})

	// add Swagger UI endpoint with hardcoded uri for simplicity
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}

func middlewareContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderKeyContentType, HeaderValueContentTypeJSON)
		next.ServeHTTP(w, r)
	})
}
