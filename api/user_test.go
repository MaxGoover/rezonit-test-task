package api

import (
	"fmt"
	"github.com/golang/mock/gomock"
	mockdb "github.com/maxgoover/rezonit-test-task/db/mock"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/util"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserAPI(t *testing.T) {
	user := randomUser()

	// Далее, нужно создать фиктивную бд
	// Для этого используем функцию NewMockStorage()
	ctrl := gomock.NewController(t)
	// Объект, который управляет состоянием теста

	// Функция Finish() проверит, все ли функции теста выполнились и в полном объеме
	defer ctrl.Finish()

	// Создаем новое хранилище
	storage := mockdb.NewMockStorage(ctrl)

	// Теперь, создаем заглушки для метода GetUser()
	// Для определения заглушки мы должны указать с какими двумя значениями аргументов эта функция GetUser()
	// будет вызываться
	storage.EXPECT().
		// Определение этой заглушки звучит так:
		// "Я ожидаю, что функция GetUser() будет вызываться с любым контекстом и конкретным аргументом
		// идентификатора учетной записи"
		GetUser(gomock.Any(), gomock.Eq(user.ID)).
		// Мы хотим, чтобы функция GetUser() вызвалась один раз
		Times(1).
		// Мы хотим, чтобы функция GetUser() возвращала объекта user и нулевую ошибку
		Return(user, nil)
	// Теперь наша заглушка готова

	// Теперь "запустим" наш сервер, но вместо реальной бд, подсунем ему фиктивную бд mockdb
	server := newTestServer(t, storage)

	// Чтобы не запускать сервер по-настоящему, мы будем использовать функцию записи пакета httptest
	// для записи ответа на запрос API
	// Для создания нового ResponseRecorder
	recorder := httptest.NewRecorder()

	// Url - который мы хотим вызвать - протестировать
	url := fmt.Sprintf("/users/%d", user.ID)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	// Отправляем запрос через маршрутизатор сервера и запишет его ответ в рекордере
	server.router.ServeHTTP(recorder, request)

	// И затем проверяем полученный ответ
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomUser() db.User {
	return db.User{
		ID:        util.GenerateRandomInt(1, 1000),
		FirstName: util.GenerateRandomString(7),
		LastName:  util.GenerateRandomString(9),
		Age:       util.GenerateRandomInt(1, 100),
	}
}
