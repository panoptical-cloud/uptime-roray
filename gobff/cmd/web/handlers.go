package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"pc-uptime/bff/db/repo"
	"pc-uptime/bff/utils"
	"strconv"
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

func (app *application) getIpFromHost(w http.ResponseWriter, r *http.Request) {
	type Reqbody struct {
		Hostname string `json:"hostname"`
	}
	var reqp Reqbody
	err := json.NewDecoder(r.Body).Decode(&reqp)
	if err != nil {
		app.logger.Error("Error decoding request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}
	ip, err := utils.GetIpFromHost(reqp.Hostname)
	if err != nil {
		app.logger.Error("Error getting IP from hostname", "error", err, "hostname", reqp.Hostname)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error resolving hostname to IP"))
		return
	}
	// Return the result as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"ip": *ip})
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

func (app *application) createServerGroup(w http.ResponseWriter, r *http.Request) {
	var reqp repo.CreateServerGroupParams
	err := json.NewDecoder(r.Body).Decode(&reqp)
	if err != nil {
		app.logger.Error("Error decoding request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}
	db, err := app.repo.CreateServerGroup(r.Context(), reqp)

	if err != nil {
		app.logger.Error("Error creating server group", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating server group"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(db)
}

func (app *application) listServerGroups(w http.ResponseWriter, r *http.Request) {
	db, err := app.repo.ListServerGroups(r.Context())
	if err != nil {
		app.logger.Error("Error listing server group", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error listing server group"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func (app *application) getServerGroup(w http.ResponseWriter, r *http.Request) {
	_groupId := r.PathValue("id")
	groupId, err := strconv.Atoi(_groupId)
	if err != nil {
		app.logger.Error("Error converting group id to int", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error converting group id to int"))
		return
	}
	db, err := app.repo.GetServerGroup(r.Context(), int64(groupId))
	if err != nil {
		app.logger.Error("Error getting server group", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting server group"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func (app *application) addServerToGroup(w http.ResponseWriter, r *http.Request) {
	var reqp repo.CreateServerParams
	err := json.NewDecoder(r.Body).Decode(&reqp)
	reqp.ID = strconv.Itoa(int(reqp.GroupID)) + "::" + reqp.Ip
	if err != nil {
		app.logger.Error("Error decoding request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}
	db, err := app.repo.CreateServer(r.Context(), reqp)
	if err != nil {
		app.logger.Error("Error adding server to group", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error adding server to group"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func (app *application) listServersByGroup(w http.ResponseWriter, r *http.Request) {
	_groupId := r.PathValue("id")
	groupId, err := strconv.Atoi(_groupId)
	if err != nil {
		app.logger.Error("Error converting group id to int", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error converting group id to int"))
		return
	}
	db, err := app.repo.ListServersByGroup(r.Context(), int64(groupId))
	if err != nil {
		app.logger.Error("Error listing servers by group", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error listing servers by group"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func (app *application) getServerById(w http.ResponseWriter, r *http.Request) {
	_groupId := r.PathValue("gid")
	serverId := r.PathValue("sid")
	groupId, err := strconv.Atoi(_groupId)
	if err != nil {
		app.logger.Error("Error converting group id to int", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error converting group id to int"))
		return
	}
	db, err := app.repo.GetServerByGidSid(r.Context(), repo.GetServerByGidSidParams{
		GroupID: int64(groupId),
		ID:      _groupId + "::" + serverId,
	})

	if err != nil {
		app.logger.Error("Error getting server by id", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting server by id"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func (app *application) generateServerToken(w http.ResponseWriter, r *http.Request) {
	groupId := r.PathValue("gid")
	serverId := r.PathValue("sid")
	// serverId, err := strconv.Atoi(_serverId)
	// if err != nil {
	// 	app.logger.Error("Error converting server id to int", "error", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Error converting server id to int"))
	// 	return
	// }

	key := make([]byte, 10)
	_, err := rand.Read(key)
	if err != nil {
		app.logger.Error("Error generating random key", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating random key"))
		return
	}
	token := fmt.Sprintf("%x", key)

	// Save the key in the database
	err = app.repo.UpdateOneTimeTokenForServerRegistration(r.Context(), repo.UpdateOneTimeTokenForServerRegistrationParams{
		ID:           groupId + "::" + serverId,
		OneTimeToken: &token,
	})
	if err != nil {
		app.logger.Error("Error saving one time token", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error saving one time token"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (app *application) verifyServerToken(w http.ResponseWriter, r *http.Request) {
	incomingToken := r.PathValue("token")
	serverId := r.PathValue("sid")
	// serverId, err := strconv.Atoi(_serverId)
	// if err != nil {
	// 	app.logger.Error("Error converting server id to int", "error", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Error converting server id to int"))
	// 	return
	// }

	db, err := app.repo.GetOneTimeTokenForServerRegistration(r.Context(), serverId)
	if err != nil {
		app.logger.Error("Error verifying token", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error verifying token"))
		return
	}

	// Verify the token
	if *db != incomingToken {
		app.logger.Error("Invalid token", "error", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid token"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	retMsg := struct {
		NatsUrl string `json:"natsUrl"`
	}{
		NatsUrl: app.natsServer.ClientURL(),
	}
	json.NewEncoder(w).Encode(retMsg)
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
