package main

import (
	"encoding/json"
	"fmt"
	"pc-uptime/bff/db/repo"
	"time"

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
		app.logger.Error("Error getting server port", "error", err)
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

func (app *application) getServerStatusEvents(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulate sending events (you can replace this with real data)
	for i := 0; i < 10000; i++ {
		fmt.Fprintf(w, "event: server-1\ndata: {\"disk\": %d, \"cpu\": %d, \"ram\": %d}\n\n", i*1, i*2, i*3)
		time.Sleep(1 * time.Second)
		w.(http.Flusher).Flush()
	}

	// Simulate closing the connection
	closeNotify := w.(http.CloseNotifier).CloseNotify()
	<-closeNotify
}
