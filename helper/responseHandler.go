package helper

import "strings"

//Response handler
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj handler
type EmptyObj struct{}

//BuildResponse return data response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse return error response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitErr := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splitErr,
		Data:    data,
	}
	return res
}
