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
		Joins("LEFT JOIN followers f ON rooms.follower_id = f.id AND f.deleted_at IS NULL").
		Where(c.db.DB.Where("f.user_id = ?", userID).Or("f.follow_user_id = ?", userID)).
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Table("(?) as users", c.db.DB.Model(&models.RoomsUser{}).
				Select("users.*, r.id AS room_id").
				Joins("JOIN followers f ON (f.follow_user_id = users.id OR f.user_id = users.id) AND (f.user_id = users.id OR f.follow_user_id = users.id)").
				Joins("JOIN rooms r ON r.follower_id = f.id"))
		})

	return userRooms, queryBuilder.Find(&userRooms).Limit(15).Error
}

func (c RoomRepository) GetRoomById(roomId int64) (room models.Room, err error) {
	return room, c.db.DB.Model(&models.Room{}).Where("id =?", roomId).First(&room).Error
}

func (c RoomRepository) GetRoomByFollowerId(followerId int64) (room models.Room, err error) {
	return room, c.db.DB.Model(&models.Room{}).
		Where("follower_id = ?", followerId).
		First(&room).
		Error
}
