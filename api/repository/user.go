package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"gorm.io/gorm"
	"time"
)

//UserRepository -> database structure
type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

//NewUserRepository -> creates a new User repository
func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

//WithTrx enables repository with transaction
func (c UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

//Create -> Create User
func (c UserRepository) Create(User models.User) (models.User, error) {
	return User, c.db.DB.Create(&User).Error
}

//GetAllUsers -> Get all users
func (c UserRepository) GetAllUsers(pagination utils.Pagination, cursor string, userId string) (users []models.UserWithFollow, count int64, err error) {
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.User{}).
		Select("users.*, (?) as follow_status",
			c.db.DB.Model(&models.Followers{}).
				Select("IF (followers.user_id IS NOT NULL, true, false)").
				Where("followers.user_id = ?", userId).
				Or("followers.follow_user_id = ?", userId).
				Limit(1)).
		Where("users.id != ?", userId)

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where("`users`.`email` LIKE ?", searchQuery)
	}

	if cursor != "" {
		time, _ := time.Parse(time.RFC3339, cursor)
		queryBuilder = queryBuilder.Where("created_at < ?", time)
	}

	return users, totalRows, queryBuilder.
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
}

//GetOneUserById -> Get one user by id
func (c UserRepository) GetOneUserById(userId string) (user models.User, err error) {
	return user, c.db.DB.Model(&user).Where("id = ?", userId).First(&user).Error
}
