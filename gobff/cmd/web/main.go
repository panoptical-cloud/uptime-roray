package main

import (
	"database/sql"

	_ "embed"
	"flag"
	"log"
	"log/slog"
	"net/http"
	repo "pc-uptime/bff/db/repo"

	_ "modernc.org/sqlite"
)

type application struct {
	logger *slog.Logger
	repo   *repo.Queries
	// sm *scs.SessionManager
	// siteUserSvc    *svc.SiteAndUserSvc
	// siteBrokerConn *nats.Conn
}

func main() {
	// ctx := context.Background()

	addr := flag.String("addr", ":9191", "HTTP network address")
	dsn := flag.String("dsn", "data.db", "DB URL")
	logL := flag.String("log", "info", "Log level (debug, info, warn, error, fatal)")
	flag.Parse()

	db, err := sql.Open("sqlite", *dsn)

	if err != nil {
		log.Fatal(err)
	}

	queries := repo.New(db)

	app := &application{
		logger: initLogger(*logL),
		repo:   queries,
		// sm: initSessionMgr("sqlite", *dsn),
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	app.logger.Info("Starting server on", slog.String("addr", *addr))
	err = srv.ListenAndServe()
	log.Fatal(err)
}
