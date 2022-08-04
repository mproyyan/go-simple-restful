package service

import (
	"context"
	"database/sql"

	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/http/request"
	"github.com/mproyyan/go-simple-restful/http/response"
	"github.com/mproyyan/go-simple-restful/model"
	"github.com/mproyyan/go-simple-restful/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
}

func NewProductService(pr repository.ProductRepository, db *sql.DB) *ProductService {
	return &ProductService{
		ProductRepository: pr,
		DB:                db,
	}
}

func (ps *ProductService) Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
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
	tx, err := ps.DB.Begin()
	helper.CheckErr(err)
	defer helper.CommitOrRollback(tx)

	data, err := ps.ProductRepository.Find(ctx, tx, request.Id)
	if err != nil {
		panic(err)
	}

	product := ps.ProductRepository.Update(ctx, tx, data)

	return helper.ProductResponse(product)
}

func (ps *ProductService) Delete(ctx context.Context, productId int) bool {
	tx, err := ps.DB.Begin()
	helper.CheckErr(err)
	defer helper.CommitOrRollback(tx)

	data, err := ps.ProductRepository.Find(ctx, tx, productId)
	if err != nil {
		panic(err)
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
		panic(err)
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
