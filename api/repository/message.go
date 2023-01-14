package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
)

//MessageRepository -> MessageRepository
type MessageRepository struct {
	db infrastructure.Database
}

//NewMessageRepository -> MessageRepository
func NewMessageRepository(
	db infrastructure.Database,
) MessageRepository {
	return MessageRepository{
		db: db,
	}
}

//GetMessagesWithUser -> Get messages with user
func (c MessageRepository) GetMessagesWithUser(roomId int64) (messages []models.Message, err error) {
	return messages, c.db.DB.Model(&models.Message{}).Where("room_id = ?", roomId).Order("created_at DESC").Find(&messages).Error
}

//SaveMessageToRoom -> Save message to room
func (c MessageRepository) SaveMessageToRoom(message models.Message) (models.Message, error) {
	return message, c.db.DB.Create(&message).Error
}
