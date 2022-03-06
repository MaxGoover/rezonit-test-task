package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type responseJSON struct {
	result string
	data   interface{}
}

func responseError(w http.ResponseWriter, err error, httpStatus int) {
	m := responseJSON{
		result: "error",
		data:   err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)
	fmt.Fprintln(w, string(res))
}

func responseOk(w http.ResponseWriter, responseData interface{}) {
	m := responseJSON{
		result: "ok",
		data:   responseData,
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(res))
}
