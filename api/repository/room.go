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

func (c RoomRepository) GetRoomWithUser(userID string, cursor string) (userRooms []models.RoomWithUsers, err error) {
	queryBuilder := c.db.DB.Model(&models.RoomWithUsers{})
	if cursor != "" {
		parsedCursor, _ := time.Parse(time.RFC3339, cursor)
		queryBuilder = queryBuilder.Where("rooms.created_at < ?", parsedCursor)
	}

	queryBuilder = queryBuilder.
		Select("rooms.*, F.follow_user_id, F.user_id").
		Joins("LEFT JOIN user_rooms ON rooms.id = user_rooms.room_id").
		Joins("LEFT JOIN followers F ON user_rooms.follower_id = F.id AND F.deleted_at IS NULL").
		Where(c.db.DB.Where("F.user_id = ?", userID).Or("F.follow_user_id = ?", userID)).
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Table("(?) as users", c.db.DB.Model(&models.RoomsUser{}).
				Select("users.*, room_id").
				Joins("JOIN followers F ON users.id = F.follow_user_id OR users.id = F.user_id").
				Joins("LEFT JOIN user_rooms ur ON ur.follower_id = F.id"))
		})

	return userRooms, queryBuilder.Find(&userRooms).Limit(15).Error
}

func (c RoomRepository) GetRoomById(roomId int64) (room models.Room, err error) {
	return room, c.db.DB.Model(&models.Room{}).Where("id =?", roomId).First(&room).Error
}
