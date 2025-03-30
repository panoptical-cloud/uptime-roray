package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetAllServerGroupsSvc(ch chan []*ServerGroup) {
	resp, err := http.Get("http://localhost:9191/api/v1/server-groups")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}
	var serverGroups []*ServerGroup
	if err := json.NewDecoder(resp.Body).Decode(&serverGroups); err != nil {
		return
	}
	ch <- serverGroups
}

func GetServersByGroupIdSvc(groupId int, ch chan []*Server) {
	resp, err := http.Get("http://localhost:9191/api/v1/server-groups/" + strconv.Itoa(groupId) + "/servers")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}
	var servers []*Server
	if err := json.NewDecoder(resp.Body).Decode(&servers); err != nil {
		log.Fatal(err)
		return
	}
	ch <- servers
}
