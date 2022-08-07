package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mproyyan/go-simple-restful/helper"
	"github.com/mproyyan/go-simple-restful/middleware"
)

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	server := http.Server{
		Addr:    "localhost:1307",
		Handler: middleware,
	}

	return &server
}

func main() {
	server := InitializeServer()

	err := server.ListenAndServe()
	helper.CheckErr(err)
}
