package main

import (
	"encoding/json"
	"net/http"
)

func GetAllServerGroupsApi(ch chan []*ServerGroup) []*ServerGroup {
	resp, err := http.Get("http://localhost:9191/api/v1/server-groups")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	var serverGroups []*ServerGroup
	if err := json.NewDecoder(resp.Body).Decode(&serverGroups); err != nil {
		return nil
	}
	return serverGroups
}

func GetServersByGroupIdApi(groupId int) []*Server {
	resp, err := http.Get("http://localhost:9191/api/v1/server-groups/{groupId}/servers" + string(groupId) + "/servers")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	var servers []*Server
	if err := json.NewDecoder(resp.Body).Decode(&servers); err != nil {
		return nil
	}
	return servers
}
