package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The error encountered a problem and cannot process your request", http.StatusInternalServerError)
		return
	}
}
