package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maxgoover/rezonit-test-task/api/response"
	db "github.com/maxgoover/rezonit-test-task/db/sqlc"
	"net/http"
	"strconv"
)

type createUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int32  `json:"age"`
}

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	arg := db.CreateUserParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}

	user, err := server.storage.CreateUser(server.ctx, arg)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	response.Ok(w, m)
}

type getUserRequest struct {
	ID int32 `uri:"id"`
}

func (server *Server) getUser(w http.ResponseWriter, r *http.Request) {
	var req getUserRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	req.ID = int32(id)
	user, err := server.storage.GetUser(server.ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Error(w, err, http.StatusNotFound)
			return
		}

		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   user,
	}

	response.Ok(w, m)
}

type listUsersRequest struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

func (server *Server) listUsers(w http.ResponseWriter, r *http.Request) {
	//var req listUsersRequest
	vars := r.URL.Query()

	limit, err := strconv.ParseInt(vars.Get("limit"), 10, 32)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	offset, err := strconv.ParseInt(vars.Get("offset"), 10, 32)
	if err != nil {
		response.Error(w, err, http.StatusBadRequest)
		return
	}

	arg := db.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	listUsers, err := server.storage.ListUsers(server.ctx, arg)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data":   listUsers,
	}

	response.Ok(w, m)
}
