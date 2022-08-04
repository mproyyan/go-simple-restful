package contract

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerContract interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Find(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindALl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
