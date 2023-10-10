package main

import (
	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/controllers/productscontroller"
	"github.com/najibuhuy/go-rest-api-gin/models"
)

func main() {
	route := gin.Default()
	models.ConnectDatabase()

	route.GET("/api/products", productscontroller.GetListProduct)
	route.GET("/api/products/:id", productscontroller.GetDetailProduct)
	route.POST("/api/products", productscontroller.CreateProduct)
	route.PUT("/api/products/:id", productscontroller.UpdateProduct)
	route.DELETE("/api/products", productscontroller.DeleteProduct)

	route.Run()

}
