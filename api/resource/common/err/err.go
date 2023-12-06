package err

import "net/http"

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

var (
	RespDBDataInsertFailure = []byte(`{"error": "db data insert failure"}`)
	RespDBDataAccessFailure = []byte(`{"error": "db data access failure"}`)
	RespDBDataUpdateFailure = []byte(`{"error": "db data update failure"}`)
	RespDBDataRemoveFailure = []byte(`{"error": "db data remove failure"}`)

	RespJSONEncodeFailure = []byte(`{"error": "json encode failure"}`)
	RespJSONDecodeFailure = []byte(`{"error": "json decode failure"}`)

	RespInvalidURLParamID = []byte(`{"error": "invalid url param-id"}`)

	RespNoDataFound = []byte(`{"error": "no data found"}`)
)

func ServerError(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(reps)
}

func BadRequest(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(reps)
}

func ValidationErrors(w http.ResponseWriter, reps []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(reps)
}
