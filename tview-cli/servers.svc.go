package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllServerGroupsSvc(ch chan []*ServerGroup) {
	// ch <- GetAllServerGroupsApi(ch)
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

func GetServersByGroupIdSvc(groupId int) []*Server {
	fmt.Println("SVC func call")
	ssDC := make(chan []*Server)
	ssDC <- GetServersByGroupIdApi(groupId)
	ss := <-ssDC
	fmt.Println("SVC func call done" + ss[0].Name)
	return ss
}

// func GetServersByGroupIdSvc(groupId int) []*Server {
// 	//replace below mock witha actual call to the server
// 	switch groupId {
// 	case 1:
// 		return MockServersA()
// 	case 2:
// 		return MockServersB()
// 	default:
// 		return nil
// 	}
// }
