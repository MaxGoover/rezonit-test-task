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

//type listAccountRequest struct {
//	PageID   int32 `form:"page_id" binding:"required,min=1"`          // Порядковый номер страницы
//	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"` // Максимальное количество записей на странице
//}
//
//func (server *Server) listAccount(ctx *gin.Context) {
//	var req listAccountRequest
//	if err := ctx.ShouldBindQuery(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.ListAccountsParams{
//		Limit:  req.PageSize,
//		Offset: (req.PageID - 1) * req.PageSize,
//	}
//
//	accounts, err := server.store.ListAccounts(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, accounts)
//}
