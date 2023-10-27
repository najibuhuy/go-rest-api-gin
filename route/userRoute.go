package route

import (
	"github.com/gin-gonic/gin"
	"github.com/najibuhuy/go-rest-api-gin/controllers"
	"github.com/najibuhuy/go-rest-api-gin/middleware"
)

func UserRoute(route *gin.Engine) {
	user := route.Group("/api/user")
	user.POST("/signup", controllers.SignUp)
	user.POST("/login", controllers.Login)
	user.GET("/test", middleware.RequireAuth, controllers.MidTest)

}
