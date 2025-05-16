package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createMovieHandler creates a new movie
// for the request to the "POST /v1/movies" endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// showMovieHandler shows movie by id
// for the request "GET /v1/movies/:id"
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// use ParamsFromContext() to try retrieve the params
	params := httprouter.ParamsFromContext(r.Context())

	// use ByName() method to get the value of 'id'
	// we're using unique positive ints for ID values
	// if the value cannot be retrieved
	// or is a non-positive int
	// we return 404 Not Found
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64) // base 10 with a bit size of 64
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// interpolate ID into placeholder response
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
