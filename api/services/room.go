package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type RoomService struct {
	repository repository.RoomRepository
}

func NewRoomService(repository repository.RoomRepository) RoomService {
	return RoomService{
		repository: repository,
	}
}

func (c RoomService) WithTrx(trxHandle *gorm.DB) RoomService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

func (c RoomService) CreateRoom(room models.Room) (models.Room, error) {
	return c.repository.CreateRoom(room)

}

func (c RoomService) GetRoomWithUser(ID string, cursor string) ([]models.Room, error) {
	return c.repository.GetRoomWithUser(ID, cursor)
}

func (c RoomService) GetRoomWithId(roomId int64) (room models.Room, err error) {
	return c.repository.GetRoomWithId(roomId)
}
