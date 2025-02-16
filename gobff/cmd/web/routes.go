package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/server-port", httpReqLogger()(app, app.createServerPort))
	mux.HandleFunc("GET /api/v1/server-port", httpReqLogger()(app, app.listServerPorts))
	mux.HandleFunc("GET /api/v1/events/server-status", httpReqLogger()(app, app.getServerStatusEvents))
	// mux.HandleFunc("GET /api/v1/server-port", httpReqLogger()(app, app.getServerPort))

	// server groups routes
	mux.HandleFunc("POST /api/v1/server-groups", httpReqLogger()(app, app.createServerGroup))
	mux.HandleFunc("GET /api/v1/server-groups", httpReqLogger()(app, app.listServerGroups))
	// add a server to a group
	mux.HandleFunc("POST /api/v1/server-groups/servers", httpReqLogger()(app, app.addServerToGroup))
	// list all servers in a group with id {id}
	mux.HandleFunc("GET /api/v1/server-groups/{id}/servers", httpReqLogger()(app, app.listServersByGroup))
	return mux
}
