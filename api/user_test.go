package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/golang/mock/gomock"
	mockdb "github.com/maxgoover/rezonit-test-task/db/mock"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/util"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type bodyTestCase map[string]interface{}

func TestCreateUserAPI(t *testing.T) {
	user := randomUser()

	testCases := []struct {
		name          string
		body          bodyTestCase
		buildStubs    func(storage *mockdb.MockStorage)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: bodyTestCase{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"age":        user.Age,
			},
			buildStubs: func(storage *mockdb.MockStorage) {
				arg := db.CreateUserParams{
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Age:       user.Age,
				}
				storage.EXPECT().
					CreateUser(gomock.Any(), arg).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: bodyTestCase{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"age":        user.Age,
			},
			buildStubs: func(storage *mockdb.MockStorage) {
				storage.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidFirstName",
			body: bodyTestCase{
				"first_name": "invalid-firstname#1",
				"last_name":  user.LastName,
				"age":        user.Age,
			},
			buildStubs: func(storage *mockdb.MockStorage) {
				storage.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidLastName",
			body: bodyTestCase{
				"first_name": user.FirstName,
				"last_name":  "invalid-lastname#2",
				"age":        user.Age,
			},
			buildStubs: func(storage *mockdb.MockStorage) {
				storage.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidAge",
			body: bodyTestCase{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"age":        -1,
			},
			buildStubs: func(storage *mockdb.MockStorage) {
				storage.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storage := mockdb.NewMockStorage(ctrl)
			tc.buildStubs(storage)

			server := newTestServer(t, storage)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser() db.User {
	return db.User{
		ID:        util.GenerateRandomInt(1, 1000),
		FirstName: util.GenerateRandomString(7),
		LastName:  util.GenerateRandomString(9),
		Age:       util.GenerateRandomInt(1, 100),
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.FirstName, gotUser.FirstName)
	require.Equal(t, user.LastName, gotUser.LastName)
	require.Equal(t, user.Age, gotUser.Age)
}
