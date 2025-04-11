package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"

	"github.com/gorilla/schema"
)

var maxBytes = 10 << 20

// SetMaxBytes sets the maximum number of bytes allowed in the request body. Default is 10485760.
func SetMaxBytes(mb int) {
	maxBytes = mb
}

// Read reads a json request body into the specified struct. The maximum read bytes is defined by `maxBytes`.

func Read[I any](w http.ResponseWriter, r *http.Request) (I, error) {
	var out I

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	if err := r.ParseForm(); err != nil {
		return out, errors.New("unable to parse form data")
	}

	decoder := schema.NewDecoder()

	if err := decoder.Decode(&out, r.Form); err != nil {
		return out, errors.New("invalid form input")
	}

	return out, nil
}

// Write writes a json response with the specified status and data.
func Write(w http.ResponseWriter, status int, data any, headers ...http.Header) {
	out, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshalling data: %s", err)
		return
	}

	if len(headers) > 0 {
		maps.Copy(w.Header(), headers[0])
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		fmt.Printf("error writing response: %s", err)
	}
}

// Error writes a json response with the specified error. A non HTTPError type will follow:
//   - Status: 400
//   - Message: "bad request"
//   - Detail: error.Error()
func Error(w http.ResponseWriter, err error) {
	Write(w, http.StatusBadRequest, err)
}
