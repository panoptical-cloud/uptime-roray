package sg

import (
	"bytes"
	"encoding/json"
	"errors"
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
		return
	}
	ch <- servers
}

func AddNewServerGroup(name string, desc string) error {
	body, err := json.Marshal(struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
	}{
		Name: name,
		Desc: desc,
	})
	if err != nil {
		return errors.New("JSON marshalling failed")
	}
	resp, err := http.Post("http://localhost:9191/api/v1/server-groups", "application/json; charset=UTF-8", bytes.NewBuffer(body))
	if resp.StatusCode != http.StatusCreated {
		return errors.New("Server group creation failed")
	}
	if err != nil {
		return errors.New("Server group creation failed")
	}
	return nil
}
