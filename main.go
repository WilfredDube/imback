package main

import (
	"github.com/WilfredDube/imback/handlers"
	"github.com/WilfredDube/imback/middleware"
	"github.com/WilfredDube/imback/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService = service.New()
	videoHandler = handlers.New(videoService)
)

func main() {
	server := gin.Default()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, videoHandler.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		videoHandler.Save(ctx)
	})

	server.Run(":4000")
}
