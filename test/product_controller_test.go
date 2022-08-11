package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mproyyan/go-simple-restful/app"
	"github.com/mproyyan/go-simple-restful/controller"
	"github.com/mproyyan/go-simple-restful/http/request"
	"github.com/mproyyan/go-simple-restful/http/response"
	mks "github.com/mproyyan/go-simple-restful/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Decode(r io.Reader, placeholder interface{}) {
	decoder := json.NewDecoder(r)
	decoder.Decode(placeholder)
}

// func Encode(w io.Writer, data interface{}) {
// 	encoder := json.NewEncoder(w)
// 	encoder.Encode(data)
// }

func TestProductControllerFindAllSuccess(t *testing.T) {
	service := mks.NewProductServiceContract(t)
	controller := controller.NewProductController(service)
	router := app.NewRouter(controller)

	expectedProducts := []response.ProductResponse{
		{Id: 1, Name: "Product Test 1"},
		{Id: 2, Name: "Product Test 2"},
	}

	service.On("FindAll", mock.Anything).Return(expectedProducts)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)

	router.ServeHTTP(rec, req)

	products := response.HttpResponse{}
	Decode(rec.Result().Body, &products)

	assert.Len(t, products.Data.([]interface{}), 2)
}

func TestProductControllerFindSuccess(t *testing.T) {
	service := mks.NewProductServiceContract(t)
	controller := controller.NewProductController(service)
	router := app.NewRouter(controller)

	expectedProduct := response.ProductResponse{Id: 1, Name: "find"}
	service.On("Find", mock.Anything, mock.AnythingOfType("int")).Return(expectedProduct)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/products/1", nil)

	router.ServeHTTP(rec, req)

	product := response.HttpResponse{}
	Decode(rec.Result().Body, &product)

	assert.Equal(t, expectedProduct.Name, product.Data.(map[string]interface{})["name"])
}

func TestProductControllerCreateSuccess(t *testing.T) {
	service := mks.NewProductServiceContract(t)
	controller := controller.NewProductController(service)
	router := app.NewRouter(controller)

	expectedProduct := response.ProductResponse{Id: 1, Name: "created"}
	service.On("Create", mock.Anything, mock.Anything).
		Return(func(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
			return response.ProductResponse{Id: 1, Name: request.Name}
		})

	buf := bytes.NewBuffer([]byte(`{"name": "created"}`))
	req := httptest.NewRequest(http.MethodPost, "/api/products", buf)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	product := response.HttpResponse{}
	Decode(rec.Result().Body, &product)

	assert.Equal(t, expectedProduct.Name, product.Data.(map[string]interface{})["name"])
}

func TestProductControllerUpdateSuccess(t *testing.T) {
	service := mks.NewProductServiceContract(t)
	controller := controller.NewProductController(service)
	router := app.NewRouter(controller)

	service.On("Update", mock.Anything, mock.Anything).
		Return(func(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
			return response.ProductResponse{Id: request.Id, Name: request.Name}
		})

	buf := bytes.NewBuffer([]byte(`{"name": "updated"}`))
	req := httptest.NewRequest(http.MethodPut, "/api/products/7", buf)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	product := response.HttpResponse{}
	Decode(rec.Result().Body, &product)

	fmt.Println(product)
	assert.Equal(t, "updated", product.Data.(map[string]interface{})["name"])
}

func TestProductControllerDeleteSuccess(t *testing.T) {
	service := mks.NewProductServiceContract(t)
	controller := controller.NewProductController(service)
	router := app.NewRouter(controller)

	service.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(true)

	req := httptest.NewRequest(http.MethodDelete, "/api/products/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	result := response.HttpResponse{}
	Decode(rec.Result().Body, &result)

	assert.Equal(t, 200, result.Code)
	assert.Equal(t, "OK", result.Status)
}
