package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/mproyyan/go-simple-restful/contract"
	"github.com/mproyyan/go-simple-restful/exception"
	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/http/request"
	"github.com/mproyyan/go-simple-restful/http/response"
	"github.com/mproyyan/go-simple-restful/model"
)

type ProductService struct {
	ProductRepository contract.ProductRepositoryContract
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(pr contract.ProductRepositoryContract, db *sql.DB, validator *validator.Validate) *ProductService {
	return &ProductService{
		ProductRepository: pr,
		DB:                db,
		Validate:          validator,
	}
}

func (ps *ProductService) Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	// validate request
	validationErr := ps.Validate.Struct(request)
	helper.CheckErr(validationErr)

	tx, err := ps.DB.Begin()
	helper.CheckErr(err)
	defer helper.CommitOrRollback(tx)

	data := model.Product{
		Name: request.Name,
	}

	product := ps.ProductRepository.Save(ctx, tx, data)

	return helper.ProductResponse(product)
}

func (ps *ProductService) Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	// validate request
	validationErr := ps.Validate.Struct(request)
	helper.CheckErr(validationErr)

	tx, err := ps.DB.Begin()
	helper.CheckErr(err)
	defer helper.CommitOrRollback(tx)

	data, err := ps.ProductRepository.Find(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// update name
	data.Name = request.Name

	product := ps.ProductRepository.Update(ctx, tx, data)

	return helper.ProductResponse(product)
}

func (ps *ProductService) Delete(ctx context.Context, productId int) bool {
	tx, err := ps.DB.Begin()
	helper.CheckErr(err)
	defer helper.CommitOrRollback(tx)

	data, err := ps.ProductRepository.Find(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	result := ps.ProductRepository.Delete(ctx, tx, data)
	return result
}

func (ps *ProductService) Find(ctx context.Context, productId int) response.ProductResponse {
	tx, err := ps.DB.Begin()
	helper.CheckErr(err)
	defer helper.CommitOrRollback(tx)

	product, err := ps.ProductRepository.Find(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ProductResponse(product)
}

func (ps *ProductService) FindAll(ctx context.Context) []response.ProductResponse {
	tx, err := ps.DB.Begin()
	helper.CheckErr(err)
	defer helper.CommitOrRollback(tx)

	products := ps.ProductRepository.FindAll(ctx, tx)

	return helper.ProductsResponse(products)
}
