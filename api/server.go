package api

import (
	"context"
	"github.com/gorilla/mux"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/util"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config     util.Config
	ctx        context.Context
	httpServer *http.Server
	router     *mux.Router
	storage    db.Storage
}

func NewServer(config util.Config, ctx context.Context, storage db.Storage) *Server {
	server := &Server{
		config:  config,
		ctx:     ctx,
		storage: storage,
	}
	return server
}

func (server *Server) Start(address string) {
	server.router = mux.NewRouter()
	server.router.HandleFunc("/users", server.listUsers).Methods("GET")
	server.router.HandleFunc("/users", server.createUser).Methods("POST")
	server.router.HandleFunc("/users/{id:[0-9]+}", server.getUser).Methods("GET")
	server.router.HandleFunc("/users/{id:[0-9]+}", server.updateUser).Methods("PUT")
	server.router.HandleFunc("/users/{id:[0-9]+}", server.deleteUser).Methods("DELETE")

	server.httpServer = &http.Server{
		Addr:    address,
		Handler: server.router,
	}

	err := server.httpServer.ListenAndServe() // запускаем сервер
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
	log.Println("Server started")

	return
}

func (server *Server) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//server.storage.Close() //закрываем соединение с БД

	defer func() {
		cancel()
	}()
	var err error
	if err = server.httpServer.Shutdown(ctxShutDown); err != nil { //выключаем сервер, с ограниченным по времени контекстом
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
