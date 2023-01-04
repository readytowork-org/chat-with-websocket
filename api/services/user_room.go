package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type UserRoomService struct {
	repository repository.UserRoomRepository
}

func NewUserRoomService(repository repository.UserRoomRepository) UserRoomService {
	return UserRoomService{
		repository: repository,
	}
}

func (c UserRoomService) WithTrx(trxHandle *gorm.DB) UserRoomService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

func (c UserRoomService) CreateUserRoom(userRoom models.UserRoom) error {
	return c.repository.CreateUserRoom(userRoom)
}
