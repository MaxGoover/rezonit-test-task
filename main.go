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
	// Загружаем конфигурацию из конфига приложения
	config, err := util.LoadConfig(".")
	// Итого, в config содержится экземпляр структуры Config, заполненный данными из переменной окружения
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Создаем контекст для работы контексто-зависимых частей системы
	_, cancel := context.WithCancel(context.Background())

	// Создаем канал для сигналов ОС
	c := make(chan os.Signal, 1)

	// В случае поступления сигнала завершения - уведомляем наш канал, бережно закрываем наше приложение
	signal.Notify(c, os.Interrupt)

	log.Println("Starting server")

	// Создаем соединение с БД и сохраним его для закрытия при остановке приложения
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	storage := db.NewStorage(conn)

	// Создаем сервер
	server := api.NewServer(config, storage)
	// В server содержится экземпляр структуры Server

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
	server.Start(config.ServerAddress)
}
