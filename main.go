package main

import (
	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/initializers"
	"github.com/najibuhuy/go-rest-api-gin/route"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.AutoMigrateModel()
}

func main() {
	router := gin.Default()
	route.ProductRoute(router)
	route.UserRoute(router)
	router.Run()

}
