package route

import (
	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/controllers"
)

func ProductRoute(route *gin.Engine) {
	product := route.Group("/api/product")
	product.GET("", controllers.GetListProduct)
	product.GET("/:id", controllers.GetDetailProduct)
	product.POST("", controllers.CreateProduct)
	product.PUT(":id", controllers.UpdateProduct)
	product.DELETE("", controllers.DeleteProduct)
}
