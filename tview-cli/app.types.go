package main

type AppViewScreens int

const (
	WELCOME AppViewScreens = iota
	SERVERS
)

type ServerScreenConfig struct {
	ServerGroups        []*ServerGroup
	Servers             []*Server
	SelectedServerGroup *ServerGroup
	SelectedServer      *Server
}

type AppStateConfig struct {
	CurrentScreen      AppViewScreens
	ServerScreenConfig *ServerScreenConfig
}

// func (asc *AppStateConfig) SetServerScreenConfig(sg *ServerScreenConfig) {
// 	if asc == nil {
// 		panic("AppStateConfig is nil")
// 	}

// 	if sg == nil {
// 		asc.serverScreenConfig = &ServerScreenConfig{}
// 	} else {
// 		asc.serverScreenConfig = sg
// 	}
// }

// func (asc *AppStateConfig) GetServerScreenConfig() *ServerScreenConfig {
// 	if asc == nil {
// 		panic("AppStateConfig is nil")
// 	}
// 	return asc.serverScreenConfig
// }

// func (asc *AppStateConfig) SetCurrentScreen(screen AppViewScreens) {
// 	if asc == nil {
// 		panic("AppStateConfig is nil")
// 	}
// 	asc.currentScreen = screen
// }

// func (asc *AppStateConfig) GetCurrentScreen() AppViewScreens {
// 	if asc == nil {
// 		panic("AppStateConfig is nil")
// 	}
// 	return asc.currentScreen
// }
