package exception

import "net/http"

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}
