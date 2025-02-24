// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repo

type Server struct {
	ID                 int64   `json:"id"`
	Name               string  `json:"name"`
	Hostname           string  `json:"hostname"`
	Ip                 *string `json:"ip"`
	AgentPort          *int64  `json:"agent_port"`
	AgentVersion       *string `json:"agent_version"`
	GroupID            int64   `json:"group_id"`
	OneTimeToken       *string `json:"one_time_token"`
	OneTimeTokenExpiry *int64  `json:"one_time_token_expiry"`
}

type ServerBaseStat struct {
	ServerID     int64    `json:"server_id"`
	Timestamp    int64    `json:"timestamp"`
	CpuUsage     *float64 `json:"cpu_usage"`
	MemoryUsage  *float64 `json:"memory_usage"`
	DiskUsage    *float64 `json:"disk_usage"`
	NetworkUsage *float64 `json:"network_usage"`
}

type ServerGroup struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ServerMetadatum struct {
	ServerID int64  `json:"server_id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type ServerPort struct {
	ServerID string `json:"server_id"`
	Port     int64  `json:"port"`
}
