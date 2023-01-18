package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"time"

	"gorm.io/gorm"
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

func (c RoomRepository) WithTrx(trxHandle *gorm.DB) RoomRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

func (c RoomRepository) CreateRoom(Room models.Room) (models.Room, error) {
	return Room, c.db.DB.Create(&Room).Error
}

func (c RoomRepository) GetRoomWithUser(userID string, cursor string) (userRooms []models.Room, err error) {

	queryBuilder := c.db.DB.Model(&models.Room{}).Joins("JOIN user_rooms on rooms.id = user_rooms.room_id").
		Where("user_rooms.user_id = ?", userID).Find(&userRooms).Limit(20)
	if cursor != "" {
		time, _ := time.Parse(time.RFC3339, cursor)
		queryBuilder = queryBuilder.Where("created_at < ?", time)
	}
	return userRooms, queryBuilder.Error
}

func (c RoomRepository) GetRoomWithId(roomId int64) (room models.Room, err error) {

	return room, c.db.DB.Model(&models.Room{}).Where("id =?", roomId).First(&room).Error
}
