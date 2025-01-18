package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"


	_ "modernc.org/sqlite"

)

func initLogger(level string) *slog.Logger {
	// Set the logger for the application
	loggingLevel := new(slog.LevelVar)
	appLogger := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     loggingLevel,
				AddSource: true,
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					if a.Key == slog.SourceKey {
						source := a.Value.Any().(*slog.Source)
						source.File = filepath.Base(source.File)
					}
					return a
				},
			},
		),
	)
	slog.SetDefault(appLogger)
	if strings.Compare(level, "debug") == 0 {
		loggingLevel.Set(slog.LevelDebug)
	}

	// we can now use the standard logger again as it uses the options we set above
	s := struct{ ID string }{ID: "123"}
	appLogger.Info("new store created", slog.String("store-id", s.ID))
	appLogger.Debug("new store created", slog.String("store-id", s.ID))

	return appLogger
}

// func initSessionMgr(dbDriverName string, dbDataSourceName string) *scs.SessionManager {
// 	// init session cache sqlite db
// 	db, err := sql.Open(dbDriverName, dbDataSourceName)
// 	utils.CheckErr(err)
// 	_, err = db.Exec(session_sql.InitSessionDB)
// 	utils.CheckErr(err)
// 	_, err = db.Exec(session_sql.InitSessionDBIndex)
// 	utils.CheckErr(err)
// 	// defer db.Close()

// 	sessionManager := scs.New()
// 	sessionManager.Store = sqlite3store.New(db)
// 	sessionManager.Lifetime = 12 * time.Hour
// 	return sessionManager
// }

// func initSiteAndUserSvc(dbDriverName string, dbDataSourceName string) *sus.SiteAndUserSvc {
// 	db, err := sql.Open(dbDriverName, dbDataSourceName)
// 	utils.CheckErr(err)
// 	// defer db.Close()
// 	return &sus.SiteAndUserSvc{
// 		SiteUserQueries: sur.New(db),
// 	}
// }
