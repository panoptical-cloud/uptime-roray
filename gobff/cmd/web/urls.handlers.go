package main

import (
	"encoding/json"
	"net/http"
	"pc-uptime/bff/db/repo"
	"strconv"
)

func (app *application) createHttpUrl(w http.ResponseWriter, r *http.Request) {
	var reqp repo.CreateHttpUrlConfigParams
	err := json.NewDecoder(r.Body).Decode(&reqp)
	if err != nil {
		app.logger.Error("error decoding request body", "error:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error decoding request body"))
		return
	}
	db, err := app.repo.CreateHttpUrlConfig(r.Context(), reqp)
	if err != nil {
		app.logger.Error("error creating http url config", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating http url config"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db)
}

func (app *application) getUrlConfigByBaseId(w http.ResponseWriter, r *http.Request) {
	baseid := r.PathValue("baseid")
	db, err := app.repo.GetHttpUrlConfigById(r.Context(), baseid)
	if err != nil {
		app.logger.Error("error getting http url config", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting http url config"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func (app *application) listHttpUrlConfigs(w http.ResponseWriter, r *http.Request) {
	limit := 10

	if r.URL.Query().Get("page") != "" {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			app.logger.Error("error parsing page query param", "error:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error parsing page query param"))
			return
		}
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * limit
		db, err := app.repo.ListHttpUrlConfigs(r.Context(), repo.ListHttpUrlConfigsParams{
			Limit:  int64(limit),
			Offset: int64(offset),
		})
		if err != nil {
			app.logger.Error("error listing http url configs", "error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error listing http url configs"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(db)
	} else {
		db, err := app.repo.ListHttpUrlConfigs(r.Context(), repo.ListHttpUrlConfigsParams{
			Limit:  int64(limit),
			Offset: 0,
		})
		if err != nil {
			app.logger.Error("error listing http url configs", "error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error listing http url configs"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(db)
	}
}
