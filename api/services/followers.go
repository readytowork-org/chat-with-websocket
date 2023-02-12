package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"

	"gorm.io/gorm"
)

type FollowersService struct {
	repository repository.FollowersRepository
}

func NewFollowersService(
	repository repository.FollowersRepository,
) FollowersService {
	return FollowersService{
		repository: repository,
	}
}

func (c FollowersService) WithTrx(trxHandle *gorm.DB) FollowersService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

func (c FollowersService) AddFollower(Follower models.Followers) (models.Followers, error) {
	return c.repository.AddFollower(Follower)
}

func (c FollowersService) UnFollower(Follower models.Followers) (models.Followers, error) {
	return c.repository.UnFollower(Follower)
}

func (c FollowersService) GetFollower(Follower models.Followers) (models.Followers, error) {
	return c.repository.GetFollower(Follower)
}
