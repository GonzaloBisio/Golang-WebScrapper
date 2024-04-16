package router

import (
	controllers "Golang/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/api/users")
	{
		userRouter.POST("", controllers.CreateUser)
		userRouter.GET("", controllers.GetUsers)
		userRouter.PUT(":id", controllers.UpdateUser)
		userRouter.DELETE(":id", controllers.DeleteUser)
	}
}
