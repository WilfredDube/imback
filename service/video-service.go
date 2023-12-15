package service

import "github.com/WilfredDube/imback/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (svc *videoService) Save(video entity.Video) entity.Video {
	svc.videos = append(svc.videos, video)

	return video
}

func (svc *videoService) FindAll() []entity.Video {
	return svc.videos
}
