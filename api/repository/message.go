package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
)

type MessageReposiotry struct {
	db infrastructure.Database
}

func NewMessageRepository(
	db infrastructure.Database,

) MessageReposiotry {
	return MessageReposiotry{
		db: db,
	}
}

func (c MessageReposiotry) GetMessagesWithUser(userID string, roomId int64) (messages []models.Message, err error) {
	return messages, c.db.DB.Model(&models.Message{}).Where("user_id = ?", userID).Where("room_id = ?", roomId).Find(&messages).Error
}
