package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// define envelope type
type envelope map[string]any

// writeJSON() is a helper to send responses
// Takes the following params:
// - destination http.ResponseWriter
// - HTTP status code to send
// - data to encode to JSON
// - header map with any additional headers we want to include
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	jsonResponse, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	jsonResponse = append(jsonResponse, '\n')
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
	return nil
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// use ParamsFromContext() to try retrieve the params
	params := httprouter.ParamsFromContext(r.Context())

	// use ByName() method to get the value of 'id'
	// we're using unique positive ints for ID values
	// if the value cannot be retrieved
	// or is a non-positive int
	// we return 404 Not Found
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64) // base 10 with a bit size of 64
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
