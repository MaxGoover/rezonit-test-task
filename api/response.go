package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseDTO struct {
	Result string      `json:"result"`
	Data   interface{} `json:"data"`
}

func responseError(w http.ResponseWriter, err error, httpStatus int) {
	res := &responseDTO{
		Result: "error",
		Data:   err.Error(),
	}

	resJSON, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)
	fmt.Fprintln(w, string(resJSON))
}

func responseOk(w http.ResponseWriter, responseData interface{}) {
	res := &responseDTO{
		Result: "ok",
		Data:   responseData,
	}

	resJSON, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(resJSON))
}
