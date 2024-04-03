package service

import "github.com/charleschinedum/golang_gin_api/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}
