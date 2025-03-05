package utils

type RegPostReqBody struct {
	Version   string `json:"agent_version"`
	Ip        string `json:"ip"`
	MachineId string `json:"mac"`
}
