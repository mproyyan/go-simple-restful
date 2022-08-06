package exception

import (
	"net/http"

	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/http/response"
)

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := response.HttpResponse{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}
