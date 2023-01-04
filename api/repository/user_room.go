package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"

	"gorm.io/gorm"
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

func (c UserRoomRepository) WithTrx(trxHandle *gorm.DB) UserRoomRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

func (c UserRoomRepository) CreateUserRoom(userRoom models.UserRoom) error {
	return c.db.DB.Create(&userRoom).Error
}
