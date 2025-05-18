package main

import (
	"fmt"
	"net/http"
)

// createMovieHandler creates a new movie
// for the request to the "POST /v1/movies" endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// showMovieHandler shows movie by id
// for the request "GET /v1/movies/:id"
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// interpolate ID into placeholder response
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
