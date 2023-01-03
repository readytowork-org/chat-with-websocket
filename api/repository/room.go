package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
)

type RoomRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewRoomRepository(db infrastructure.Database,
	logger infrastructure.Logger) RoomRepository {
	return RoomRepository{
		db:     db,
		logger: logger,
	}
}

func (c RoomRepository) CreateRoom(Room models.Room) error {
	return c.db.DB.Create(&Room).Error
}
