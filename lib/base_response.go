package lib

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code string `json:"code"`
	Data interface{} `json:"data"`
}

type ValidateError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ResponseError struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Errors []*ValidateError `json:"errors"`
}

func BaseResponse(data interface{}, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	j, _ := json.Marshal(data)
	w.Write(j)
	return
}
