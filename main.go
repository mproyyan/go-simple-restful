package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mproyyan/go-simple-restful/app"
	"github.com/mproyyan/go-simple-restful/controller"
	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/repository"
	"github.com/mproyyan/go-simple-restful/service"
)

func main() {
	database := app.NewDB()
	validator := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, database, validator)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(productController)

	server := http.Server{
		Addr:    "localhost:1307",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.CheckErr(err)
}
