package handlers

import (
	"github.com/WilfredDube/imback/entity"
	"github.com/WilfredDube/imback/service"
	"github.com/gin-gonic/gin"
)

type VideoHandler interface {
	Save(ctx *gin.Context)
	FindAll() []entity.Video
}

type handler struct {
	service service.VideoService
}

func New(service service.VideoService) VideoHandler {
	return &handler{
		service: service,
	}
}

func (h *handler) Save(ctx *gin.Context) {
	var video entity.Video
	ctx.BindJSON(&video)

	h.service.Save(video)

	ctx.JSON(200, gin.H{
		"video": video,
	})
}

func (h handler) FindAll() []entity.Video {
	return h.service.FindAll()
}
