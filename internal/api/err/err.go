package err

import "net/http"

var (
	// device error responses
	DeviceInUseErrResp         = []byte(`{"error": "operation cannot be completed because the device is in use"}`)
	DeviceServiceFailedErrResp = []byte(`{"error": "device operation failed"}`)
	DeviceNotFoundErrResp      = []byte(`{"error": "device not found"}`)

	// handler error responses
	JSONEncodeErrResp = []byte(`{"error": "error encoding json"}`)
	JSONDecodeErrResp = []byte(`{"error": "error decoding json"}`)
	InvalidIDErrResp  = []byte(`{"error": "invalid id param in url"}`)
)

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

func ServerError(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(error)
}

func BadRequest(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(error)
}

func NotFound(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusNotFound)
	w.Write(error)
}

func UnprocessableEntity(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(reps)
}
