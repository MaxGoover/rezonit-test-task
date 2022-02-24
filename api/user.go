package api

import (
	"net/http"
)

type createUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
}

func (server *Server) createUser(w http.ResponseWriter, r *http.Request) {
	// var req createUserRequest
	// нужно как то брать из контекста данные из реквеста
	// в gin мы делали это с помощью ctx.ShouldBindJSON(&req)

	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))

	//if err := ctx.ShouldBindJSON(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	//	return
	//}
	//
	//arg := db.CreateAccountParams{
	//	Owner:    req.Owner,
	//	Currency: req.Currency,
	//	Balance:  0,
	//}
	//
	//account, err := server.store.CreateAccount(ctx, arg)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}
	//
	//ctx.JSON(http.StatusOK, account)
}

//type getAccountRequest struct {
//	ID int64 `uri:"id" binding:"required,min=1"`
//}
//
//func (server *Server) getAccount(ctx *gin.Context) {
//	var req getAccountRequest
//	if err := ctx.ShouldBindUri(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	account, err := server.store.GetAccount(ctx, req.ID)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, account)
//}
//
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
