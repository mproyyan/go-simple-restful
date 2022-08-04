package helper

import (
	"github.com/mproyyan/go-simple-restful/http/response"
	"github.com/mproyyan/go-simple-restful/model"
)

func ProductResponse(product model.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:   product.Id,
		Name: product.Name,
	}
}

func ProductsResponse(products []model.Product) []response.ProductResponse {
	var data []response.ProductResponse

	for _, product := range products {
		data = append(data, ProductResponse(product))
	}

	return data
}
