package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type UserRoomService struct {
	repository repository.UserRoomRepository
}

func NewUserRoomService(repository repository.UserRoomRepository) UserRoomService {
	return UserRoomService{
		repository: repository,
	}
}

func (c UserRoomService) CreateUserRoom(userRoom models.UserRoom) error {
	return c.repository.CreateUserRoom(userRoom)
}
