package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
)

type UserRoomRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewUserRoomRepository(db infrastructure.Database,
	logger infrastructure.Logger) UserRoomRepository {
	return UserRoomRepository{
		db:     db,
		logger: logger,
	}
}

func (c UserRoomRepository) CreateUserRoom(userRoom models.UserRoom) error {
	return c.db.DB.Create(&userRoom).Error
}
