package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type RoomService struct {
	repository repository.RoomRepository
}

func NewRoomService(repository repository.RoomRepository) RoomService {
	return RoomService{
		repository: repository,
	}
}

func (c RoomService) CreateRoom(room models.Room) error {
	err := c.repository.CreateRoom(room)
	return err
}
