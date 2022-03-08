package api

import (
	"context"
	"github.com/gorilla/mux"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/util"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	config   util.Config
	ctx      context.Context
	errorLog *log.Logger
	infoLog  *log.Logger
	router   *mux.Router
	srv      *http.Server
	storage  db.Storage
}

func NewServer(config util.Config, ctx context.Context, storage db.Storage) *Server {
	server := &Server{
		config:  config,
		ctx:     ctx,
		storage: storage,
	}
	return server
}

func (server *Server) setupRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.listUsers).Methods("GET")
	router.HandleFunc("/users", server.createUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", server.getUser).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", server.updateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", server.deleteUser).Methods("DELETE")

	server.router = router
}

func (server *Server) Start() {
	server.setupRouter()
	server.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	server.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	server.srv = &http.Server{
		Addr:         server.config.ServerAddress,
		Handler:      server.router,
		ReadTimeout:  server.config.AppReadTimeout * time.Second,
		WriteTimeout: server.config.AppWriteTimeout * time.Second,
	}

	err := server.srv.ListenAndServe() // запускаем сервер
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
	log.Println("Server started")

	return
}

func (server *Server) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
