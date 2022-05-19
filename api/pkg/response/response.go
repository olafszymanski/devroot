package response

import (
	"encoding/json"
)

type IResponse interface {
	ToJSON() string
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSuccessResponse(code int, msg string) SuccessResponse {
	return SuccessResponse{
		Code:    code,
		Message: msg,
	}
}

func (r SuccessResponse) ToJSON() string {
	conv, _ := json.Marshal(r)
	return string(conv)
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func NewErrorResponse(code int, err error) ErrorResponse {
	return ErrorResponse{
		Code:  code,
		Error: err.Error(),
	}
}

func (r ErrorResponse) ToJSON() string {
	conv, _ := json.Marshal(r)
	return string(conv)
}
