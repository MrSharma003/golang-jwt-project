package routes

import (
	controller "github.com/MrSharma003/GO-JWT/controllers"
	middleware "github.com/MrSharma003/GO-JWT/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.Use(middleware.Autheticates())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUserById())

}
