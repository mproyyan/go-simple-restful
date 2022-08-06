package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/http/response"
)

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := response.HttpResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, response)
		return true
	} else {
		return false
	}
}
