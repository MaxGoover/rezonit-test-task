package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"github.com/maxgoover/rezonit-test-task/pkg/logging"
	"net/http"
)

type createUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int32  `json:"age"`
}

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	logging.Info.Println("get params from request")
	var req createUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseError(w, err, http.StatusBadRequest)
		return
	}

	arg := db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}

	logging.Info.Println("create user")
	user, err := server.storage.CreateUser(server.ctx, arg)
	if err != nil {
		responseError(w, err, http.StatusInternalServerError)
		return
	}

	responseOk(w, user)
}

type deleteUserRequest struct {
	ID int32 `uri:"id"`
}

func (server *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	logging.Info.Println("get params from request")
	var req deleteUserRequest
	paramsURL := mux.Vars(r)
	_, err := fmt.Sscan(paramsURL["id"], &req.ID)
	if err != nil {
		responseError(w, err, http.StatusBadRequest)
		return
	}

	logging.Info.Println("delete user")
	err = server.storage.DeleteUser(server.ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			responseError(w, err, http.StatusNotFound)
			return
		}

		responseError(w, err, http.StatusInternalServerError)
		return
	}

	responseOk(w, "User deleted")
}

type getUserRequest struct {
	ID int32 `uri:"id"`
}

func (server *Server) getUser(w http.ResponseWriter, r *http.Request) {
	logging.Info.Println("get params from request")
	var req getUserRequest
	paramsURL := mux.Vars(r)
	_, err := fmt.Sscan(paramsURL["id"], &req.ID)
	if err != nil {
		responseError(w, err, http.StatusBadRequest)
		return
	}

	logging.Info.Println("get user")
	user, err := server.storage.GetUser(server.ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			responseError(w, err, http.StatusNotFound)
			return
		}

		responseError(w, err, http.StatusInternalServerError)
		return
	}

	responseOk(w, user)
}

type listUsersRequest struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

func (server *Server) listUsers(w http.ResponseWriter, r *http.Request) {
	logging.Info.Println("get params from request")
	var req listUsersRequest
	vars := r.URL.Query()

	_, err := fmt.Sscan(vars.Get("limit"), &req.Limit)
	if err != nil {
		responseError(w, err, http.StatusBadRequest)
		return
	}

	_, err = fmt.Sscan(vars.Get("offset"), &req.Offset)
	if err != nil {
		responseError(w, err, http.StatusBadRequest)
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	logging.Info.Println("get list users")
	listUsers, err := server.storage.ListUsers(server.ctx, arg)
	if err != nil {
		responseError(w, err, http.StatusInternalServerError)
		return
	}

	responseOk(w, listUsers)
}

type updateUserRequest struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int32  `json:"age"`
}

func (server *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	logging.Info.Println("get params from request")
	var req updateUserRequest
	paramsURL := mux.Vars(r)
	_, err := fmt.Sscan(paramsURL["id"], &req.ID)
	if err != nil {
		responseError(w, err, http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseError(w, err, http.StatusBadRequest)
		return
	}

	arg := db.UpdateUserParams{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}

	logging.Info.Println("update user")
	user, err := server.storage.UpdateUser(server.ctx, arg)
	if err != nil {
		responseError(w, err, http.StatusInternalServerError)
		return
	}

	responseOk(w, user)
}
