package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type MessageService struct {
	repository repository.MessageReposiotry
}

func NewMessageService(
	repository repository.MessageReposiotry,
) MessageService {
	return MessageService{
		repository: repository,
	}
}

func (c MessageService) CreateMessageWithUser(roomId int64, Messages models.Message) error {
	return c.repository.CreateMessageWithUser(roomId, Messages)
}

func (c MessageService) GetMessageWithUser(roomId int64, cursor string) (messages []models.Message, err error) {
	return c.repository.GetMessagesWithUser(roomId, cursor)
}
