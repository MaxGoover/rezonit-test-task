package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/maxgoover/rezonit-test-task/api"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/util"
	"log"
	"os"
	"os/signal"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	conn, err := sql.Open(config.DBDriver, config.DBSource())
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	storage := db.NewStorage(conn)
	log.Println("Starting server...")
	server := api.NewServer(config, ctx, storage)

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		server.Shutdown()
		cancel()
	}()

	server.Start()
}
