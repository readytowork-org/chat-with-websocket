package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"

	"gorm.io/gorm"
)

//UserService -> struct
type UserService struct {
	repository repository.UserRepository
}

//NewUserService -> creates a new Userservice
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

//WithTrx -> enables repository with transaction
func (c UserService) WithTrx(trxHandle *gorm.DB) UserService {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

//CreateUser -> call to create the User
func (c UserService) CreateUser(user models.User) (models.User, error) {
	return c.repository.Create(user)

}

//GetAllUsers -> call to get all the User
func (c UserService) GetAllUsers(pagination utils.Pagination, cursor string, userId string) ([]models.UserWithFollow, int64, error) {
	return c.repository.GetAllUsers(pagination, cursor, userId)
}

//GetOneUserById -> Get one user by id
func (c UserService) GetOneUserById(userId string) (user models.User, err error) {
	return c.repository.GetOneUserById(userId)
}
