package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/maxgoover/rezonit-test-task/api"
	"github.com/maxgoover/rezonit-test-task/util"
	"log"
	"os"
	"os/signal"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/rezonit_test_task?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// Загружаем конфигурацию из конфига приложения
	config, err := util.LoadConfig(".")
	// Итого, в config содержится экземпляр структуры Config, заполненный данными из переменной окружения
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Создаем контекст для работы контексто-зависимых частей системы
	ctx, cancel := context.WithCancel(context.Background())

	// Создаем канал для сигналов ОС
	c := make(chan os.Signal, 1)

	// В случае поступления сигнала завершения - уведомляем наш канал, бережно закрываем наше приложение
	signal.Notify(c, os.Interrupt)

	// Создаем сервер
	server := api.NewServer(config, ctx)
	// В server содержится экземпляр структуры Server
	// Этот server принимает значение переменной окружения

	// Горутина для ловли сообщений системы
	go func() {
		// Если в операционную систему пришел какой-то сигнал
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		// Останавливаем сервер
		server.Shutdown()
		// Отменяем контекст
		cancel()
	}()

	// Запускаем сервер
	server.Start()

	//conn, err := sql.Open(dbDriver, dbSource)
	//if err != nil {
	//	log.Fatal("cannot connect to db:", err)
	//}
	//
	//storage := db.NewStorage(conn)
	//server := api.NewServer(storage)
	//
	//err = server.Start(serverAddress)
	//if err != nil {
	//	log.Fatal("cannot start server:", err)
	//}
}
