package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"time"
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

func (c MessageReposiotry) GetMessagesWithUser( roomId int64, cursor string) (messages []models.Message, err error) {
	queryBuilder := c.db.DB.Model(&models.Message{}).Order("created_at desc").Where("room_id =?", roomId).Find(&messages).Limit(20)
	if cursor != "" {
		time, _ := time.Parse(time.RFC3339, cursor)
		queryBuilder = queryBuilder.Where("created_at < ?", time)
	}

	return messages, queryBuilder.Error
}
