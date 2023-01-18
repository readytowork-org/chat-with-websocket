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

func (c MessageService) GetMessageWithUser(roomId int64, cursor string) (messages []models.Message, err error) {
	return c.repository.GetMessagesWithUser(roomId, cursor)
}
