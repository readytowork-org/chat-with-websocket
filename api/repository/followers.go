package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FollowersRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewFollowersRepository(
	db infrastructure.Database,
	logger infrastructure.Logger,
) FollowersRepository {
	return FollowersRepository{
		db:     db,
		logger: logger,
	}
}

func (c FollowersRepository) WithTrx(trxHandle *gorm.DB) FollowersRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

func (c FollowersRepository) AddFollower(Follower models.Followers) (models.Followers, error) {
	return Follower, c.db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&Follower).Error
}

func (c FollowersRepository) UnFollower(Follower models.Followers) (models.Followers, error) {
	return Follower, c.db.DB.
		Where("user_id = ?", Follower.UserId).
		Where("follow_user_id = ?", Follower.FollowUserId).
		Delete(&Follower).
		Error
}
