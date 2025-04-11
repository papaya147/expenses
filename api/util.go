package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
)

var maxBytes = 10 << 20

// SetMaxBytes sets the maximum number of bytes allowed in the request body. Default is 10485760.
func SetMaxBytes(mb int) {
	maxBytes = mb
}

// Read reads a json request body into the specified struct. The maximum read bytes is defined by `maxBytes`.
func Read[I any](w http.ResponseWriter, r *http.Request) (I, error) {
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	var out I
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&out); err != nil {
		return out, errors.New("bad request")
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return out, errors.New("json body must contain only one object")
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
