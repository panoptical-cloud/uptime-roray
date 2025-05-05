package utils

type RegPostReqBody struct {
	Version   string `json:"agent_version"`
	Ip        string `json:"ip"`
	MachineId string `json:"mac"`
}

type InitConf struct {
	NatsUrl     string `json:"nats_url"`
	NatsSubject string `json:"subj"`
}

type SvcMonitorConf struct {
	Path      string
	Port      int32
	DoMonitor bool
}

type AppHolder struct {
	AppConf     *InitConf
	MonitorSvcs []*SvcMonitorConf
}
