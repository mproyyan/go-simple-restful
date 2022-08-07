//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/mproyyan/go-simple-restful/app"
	"github.com/mproyyan/go-simple-restful/contract"
	"github.com/mproyyan/go-simple-restful/controller"
	"github.com/mproyyan/go-simple-restful/middleware"
	"github.com/mproyyan/go-simple-restful/repository"
	"github.com/mproyyan/go-simple-restful/service"
)

var productRepositorySet = wire.NewSet(
	repository.NewProductRepository,
	wire.Bind(new(contract.ProductRepositoryContract), new(*repository.ProductRepository)),
)

var productServiceSet = wire.NewSet(
	service.NewProductService,
	wire.Bind(new(contract.ProductServiceContract), new(*service.ProductService)),
)

var productControllerSet = wire.NewSet(
	controller.NewProductController,
	wire.Bind(new(contract.ProductControllerContract), new(*controller.ProductController)),
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		productRepositorySet,
		productServiceSet,
		productControllerSet,
		app.NewRouter,
		middleware.NewAuthMiddleware,
		wire.Bind(new(http.Handler), new(*httprouter.Router)), // NewAuthMiddleware require http.Handler iface so must be bind first
		NewServer,
	)

	return nil
}
