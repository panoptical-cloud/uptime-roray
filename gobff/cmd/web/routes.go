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

	// START: utility routes
	mux.HandleFunc("POST /api/v1/utils/ip-by-host", httpReqLogger()(app, app.getIpFromHost))
	// END: utility routes

	// START: server groups routes
	// Add new server group
	mux.HandleFunc("POST /api/v1/server-groups", httpReqLogger()(app, app.createServerGroup))

	// List all server groups
	mux.HandleFunc("GET /api/v1/server-groups", httpReqLogger()(app, app.listServerGroups))

	// Get server group by id
	mux.HandleFunc("GET /api/v1/server-groups/{id}", httpReqLogger()(app, app.getServerGroup))

	// add a server to a group
	mux.HandleFunc("POST /api/v1/server-groups/servers", httpReqLogger()(app, app.addServerToGroup))

	// list all servers in a group with id {id}
	mux.HandleFunc("GET /api/v1/server-groups/{id}/servers", httpReqLogger()(app, app.listServersByGroup))

	// get server by group id and server id
	mux.HandleFunc("GET /api/v1/server-groups/{gid}/servers/{sid}", httpReqLogger()(app, app.getServerById))

	// generate one time registration token for a new server
	mux.HandleFunc("GET /api/v1/server-groups/{gid}/servers/{sid}/regtoken", httpReqLogger()(app, app.generateServerToken))

	// verify server registration token
	mux.HandleFunc("POST /api/v1/server/{sid}/verifytoken/{token}", httpReqLogger()(app, app.verifyServerToken))

	// END: server groups routes

	return mux
}
