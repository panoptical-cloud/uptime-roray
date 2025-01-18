package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/server-port", httpReqLogger()(app, app.createServerPort))
	mux.HandleFunc("GET /api/v1/server-port", httpReqLogger()(app, app.listServerPorts))
	// mux.HandleFunc("GET /api/v1/server-port", httpReqLogger()(app, app.getServerPort))

	return mux
}
