package main

import (
	"database/sql"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	_ "embed"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"pc-uptime/bff/api"
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

	nc, err := nats.Connect("nats://199.241.138.81:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	log.Println("Connected to 199.241.138.81:4222")

	nc.Subscribe("test.rpc", func(msg *nats.Msg) {

		rcvData := &api.BaseStatsReply{}
		err = proto.Unmarshal(msg.Data, rcvData)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received RPC request: %+v\n", rcvData)

		// repData := &RpcReply{}
		// repData.Txt = fmt.Sprintf("This is a response %v, from count %v", uniqueID, rcvData.Txt)

		// repMsg, err := proto.Marshal(repData)
		// checkErr(err)
		// err = nc.Publish(msg.Reply, repMsg)
		// checkErr(err)
	})

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
