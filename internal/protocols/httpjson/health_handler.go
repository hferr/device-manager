package httpjson

import "net/http"

// @Summary      Health check
// @Description  Endpoint to perform a health check on the system
// @Tags         Health
// @Success      200
// @Router       /health [get]
func (h Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
