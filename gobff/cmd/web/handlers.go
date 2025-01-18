package main

import (
	"encoding/json"
	"pc-uptime/bff/db/repo"

	"net/http"
)

func (app *application) createServerPort(w http.ResponseWriter, r *http.Request) {
	var reqp repo.CreateServerPortParams
	err := json.NewDecoder(r.Body).Decode(&reqp)
	if err != nil {
		app.logger.Error("Error decoding request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}
	db, err := app.repo.CreateServerPort(r.Context(), reqp)

	if err != nil {
		app.logger.Error("Error creating server port", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating server port"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db)
}

func (app *application) getServerPort(w http.ResponseWriter, r *http.Request) {
	var reqp repo.GetServerPortParams
	err := json.NewDecoder(r.Body).Decode(&reqp)
	if err != nil {
		app.logger.Error("Error decoding request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}
	db, err := app.repo.GetServerPort(r.Context(), reqp)
	if err != nil {
		app.logger.Error("Error getting server port",  "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting server port"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func (app *application) listServerPorts(w http.ResponseWriter, r *http.Request) {
	db, err := app.repo.ListServerPorts(r.Context())
	if err != nil {
		app.logger.Error("Error listing server port", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error listing server port"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}
