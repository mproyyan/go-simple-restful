package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mproyyan/go-simple-restful/contract"
	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/http/request"
	"github.com/mproyyan/go-simple-restful/http/response"
	"github.com/mproyyan/go-simple-restful/service"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(ps service.ProductService) contract.ProductControllerContract {
	return &ProductController{
		ProductService: ps,
	}
}

func (pc *ProductController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productCreateRequest := request.ProductCreateRequest{}
	helper.ReadFromRequestBody(r, &productCreateRequest)

	result := pc.ProductService.Create(r.Context(), productCreateRequest)
	response := response.HttpResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   result,
	}

	helper.WriteToResponseBody(w, response)
}

func (pc *ProductController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productUpdateRequest := request.ProductUpdateRequest{}
	helper.ReadFromRequestBody(r, &productUpdateRequest)

	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.CheckErr(err)

	productUpdateRequest.Id = id

	result := pc.ProductService.Update(r.Context(), productUpdateRequest)
	response := response.HttpResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   result,
	}

	helper.WriteToResponseBody(w, response)
}

func (pc *ProductController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.CheckErr(err)

	pc.ProductService.Delete(r.Context(), id)
	response := response.HttpResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, response)
}

func (pc *ProductController) Find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.CheckErr(err)

	result := pc.ProductService.Find(r.Context(), id)
	response := response.HttpResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   result,
	}

	helper.WriteToResponseBody(w, response)
}

func (pc *ProductController) FindALl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := pc.ProductService.FindAll(r.Context())
	response := response.HttpResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   result,
	}

	helper.WriteToResponseBody(w, response)
}
