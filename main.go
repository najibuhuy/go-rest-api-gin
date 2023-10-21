package main

import (
	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/models"
	"github.com/najibuhuy/go-rest-api-gin/route"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	route.ProductRoute(router)

	router.Run()

}
