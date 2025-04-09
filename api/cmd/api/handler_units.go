package main

import "net/http"

func (app *application) listUnitsHandler(w http.ResponseWriter, r *http.Request) {
	units, err := app.repos.Units.List()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := envelope{"data": map[string]any{
		"units": units,
	}}

	if err = app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
