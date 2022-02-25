package api

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/maxgoover/rezonit-test-task/api/middleware"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/util"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config  util.Config
	ctx     context.Context
	storage *db.Storage
	srv     *http.Server
}

func NewServer(config util.Config, ctx context.Context) *Server {
	server := &Server{
		ctx:    ctx,
		config: config,
	}

	return server
}

func (server *Server) Start() {
	log.Println("Starting server")
	log.Println(server.config.GetDBString())

	// Создаем соединение с БД и сохраним его для закрытия при остановке приложения
	conn, err := sql.Open(server.config.DBDriver, server.config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	server.storage = db.NewStorage(conn)

	//carsStorage := db3.NewCarStorage(server.db)    //создаем экземпляр storage для работы с бд и всем что связано с машинами
	//usersStorage := db3.NewUsersStorage(server.storage) //создаем экземпляр storage для работы с бд и всем что связано с пользователями

	//carsProcessor := processors.NewCarsProcessor(carsStorage) //инициализируем процессоры соотвествующими storage
	//usersProcessor := processors.NewUsersProcessor(usersStorage)

	//userHandler := handlers.NewUsersHandler(usersProcessor) //инициализируем handlerы нашими процессорами
	//carsHandler := handlers.NewCarsHandler(carsProcessor)

	routes := mux.NewRouter()
	routes.HandleFunc("/users", server.createUser).Methods("POST")
	routes.HandleFunc("/users/{id:[0-9]+}", server.getUser).Methods("GET")
	//routes.HandleFunc("/users", server.ListUser).Methods("GET")
	//routes.HandleFunc("/users/list", userHandler.List).Methods("GET")
	//router.POST("/users", server.createUser)
	//router.GET("/users/:id", server.getUser)
	//router.GET("/users", server.listUsers)

	// Используем middleware посредников
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    server.config.ServerAddress,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe() // запускаем сервер

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

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
	if err = server.srv.Shutdown(ctxShutDown); err != nil { //выключаем сервер, с ограниченным по времени контекстом
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
