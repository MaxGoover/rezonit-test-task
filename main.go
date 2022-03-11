package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/maxgoover/rezonit-test-task/api"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/pkg/logging"
	"github.com/maxgoover/rezonit-test-task/util"
	"os"
	"os/signal"
)

func main() {
	logging.Info.Println("load config")
	config, err := util.LoadConfig(".")
	if err != nil {
		logging.Error.Fatal("cannot load config:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	logging.Info.Println("connect to db")
	conn, err := sql.Open(config.DBDriver, config.DBSource())
	if err != nil {
		logging.Error.Fatal("cannot connect to db:", err)
	}

	logging.Info.Println("start server")
	storage := db.NewStorage(conn)
	server := api.NewServer(ctx, config, storage)

	go func() {
		oscall := <-c
		logging.Info.Printf("system call:%+v", oscall)
		server.Shutdown()
		cancel()
	}()

	server.Start()
}
