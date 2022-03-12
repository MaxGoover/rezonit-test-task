package api

import (
	"context"
	"github.com/gorilla/mux"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/pkg/logging"
	"github.com/maxgoover/rezonit-test-task/util"
	"net/http"
	"time"
)

type Server struct {
	config  util.Config
	ctx     context.Context
	router  *mux.Router
	srv     *http.Server
	storage db.Storage
}

func NewServer(ctx context.Context, config util.Config, storage db.Storage) *Server {
	server := &Server{
		config:  config,
		ctx:     ctx,
		storage: storage,
	}
	return server
}

func (server *Server) setupRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users", server.createUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/users", server.listUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/{id:[0-9]+}", server.getUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/{id:[0-9]+}", server.updateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/users/{id:[0-9]+}", server.deleteUser).Methods("DELETE", "OPTIONS")

	server.router = router
}

func (server *Server) Start() {
	server.setupRouter()
	server.srv = &http.Server{
		Addr:    server.config.ServerAddress,
		Handler: server.router,
	}

	logging.Info.Println("server started...")
	err := server.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logging.Error.Fatal(err)
	}

	return
}

func (server *Server) Shutdown() {
	logging.Info.Printf("server stops")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		logging.Error.Fatalf("server Shutdown Failed:%s", err)
	}

	logging.Info.Printf("server exited properly")
	if err == http.ErrServerClosed {
		err = nil
	}
}
