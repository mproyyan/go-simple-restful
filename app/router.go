package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mproyyan/go-simple-restful/controller"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/products", productController.FindALl)
	router.GET("/api/products/:productId", productController.Find)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	return router
}
