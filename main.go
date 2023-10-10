package main

import (
	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/controllers/productcontroller"
	"github.com/najibuhuy/go-rest-api-gin/model"
)

func main() {
	route := gin.Default()
	model.ConnectDatabase()

	route.GET("/api/products", productcontroller.GetListProduct)
	route.GET("/api/products/:id", productcontroller.GetDetailProduct)
	route.POST("/api/products", productcontroller.CreateProduct)
	route.PUT("/api/products/:id", productcontroller.UpdateProduct)
	route.DELETE("/api/products", productcontroller.DeleteProduct)

	route.Run()

}
