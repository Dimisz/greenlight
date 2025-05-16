package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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
