package main

import (
	"fmt"
	"net/http"
)

// logError() is generic helper for logging error messages
// TODO: upgrade later
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// errorResponse() is a generic helper for sending
// JSON-formatted error messages to the client
// with a given status code
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse() is used when our app encounters
// an unexpected problem at runtime. It logs the detailed
// error message then uses errorResponse() helper to send a 500
// Internal Server Error and JSON response
// with generic error message
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse() is used to send 404 Not Found status code
// and JSON response to the client
// satisfies http.Handler interface
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowed() is used to send 405 Method Not Allowed
// status code and JSON response to the client
// satisfies http.Handler interface
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
