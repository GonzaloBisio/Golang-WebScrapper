package router

import (
	league_controller "Golang/Controllers/football.controller/league.controller"
	football_controller "Golang/Controllers/football.controller/player.controller"
	message_controller "Golang/Controllers/message.controller"
	user_controller "Golang/Controllers/user.controller"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

func Setup(router *gin.Engine) {
	c := cors.Default().Handler(router)

	userRouter := router.Group("/api/users")
	{
		userRouter.POST("", user_controller.CreateUser)
		userRouter.GET("", user_controller.GetUsers)
		userRouter.PUT(":id", user_controller.UpdateUser)
		userRouter.DELETE(":id", user_controller.DeleteUser)
	}
	userRouter = router.Group("/api/football")
	{
		userRouter.GET("player", football_controller.GetPalyerStats)
		userRouter.GET("league", league_controller.GetLeagueStats)
	}
	userRouter = router.Group("/api/messages")
	{
		userRouter.POST("", message_controller.SendMessage)
	}

	// Start the HTTP server with the CORS handler
	http.ListenAndServe(":8080", c)
}
