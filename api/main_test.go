package api

import (
	"context"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/util"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func newTestServer(t *testing.T, store db.Storage) *Server {
	// Загружаем конфигурацию из конфига приложения
	config, err := util.LoadConfig("..")
	// Итого, в config содержится экземпляр структуры Config, заполненный данными из переменной окружения
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := NewServer(config, context.Background(), store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
