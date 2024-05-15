package main

import (
	"Golang/router"
	"github.com/gin-gonic/gin"
)

// @title Tag Service Api
// @version 1.0
// @description A tag service with swagger integration

// @host localhost:8888
// @BasePath /api

func main() {
	r := gin.Default() // Create a new gin.Engine instance
	router.Setup(r)    // Pass the gin.Engine instance to SetupUserRoutes
}
