package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type MessageService struct {
	repository repository.MessageRepository
}

func NewMessageService(
	repository repository.MessageRepository,
) MessageService {
	return MessageService{
		repository: repository,
	}
}

func (c MessageService) GetMessageWithUser(roomId int64) (messages []models.Message, err error) {
	return c.repository.GetMessagesWithUser(roomId)
}

//SaveMessageToRoom -> Save message to room
func (c MessageService) SaveMessageToRoom(message models.Message) (models.Message, error) {
	return c.repository.SaveMessageToRoom(message)
}
