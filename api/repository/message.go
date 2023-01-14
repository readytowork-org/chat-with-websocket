package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
)

//MessageRepository -> MessageRepository
type MessageRepository struct {
	db             infrastructure.Database
	userRepository UserRepository
}

//NewMessageRepository -> MessageRepository
func NewMessageRepository(
	db infrastructure.Database,
	userRepository UserRepository,
) MessageRepository {
	return MessageRepository{
		db:             db,
		userRepository: userRepository,
	}
}

//GetMessagesWithUser -> Get messages with user
func (c MessageRepository) GetMessagesWithUser(roomId int64) (messages []models.UserMessage, err error) {
	return messages, c.db.DB.Model(&messages).Where("room_id = ?", roomId).Order("created_at DESC").Preload("User").Find(&messages).Error
}

//SaveMessageToRoom -> Save message to room
func (c MessageRepository) SaveMessageToRoom(message models.UserMessage) (models.UserMessage, error) {
	err := c.db.DB.Create(&message).Error
	if err != nil {
		return message, err
	}
	user, err := c.userRepository.GetOneUserById(message.UserId)
	if err != nil {
		return message, err
	}

	message.User = user

	return message, err
}
