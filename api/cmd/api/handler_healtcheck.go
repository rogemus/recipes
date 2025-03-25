package main

import "net/http"

func (app *application) healtcheckHandler(w http.ResponseWriter, r *http.Request) {
	response := envelope{
		"data": map[string]any{
			"status": "available",
			"system_info": map[string]string{
				"environment": app.config.env,
				"version":     version,
			}},
	}

	err := app.writeJSON(w, http.StatusOK, response, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
