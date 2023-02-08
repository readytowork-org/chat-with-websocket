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
func (c UserRepository) GetAllUsers(pagination utils.Pagination, cursor string, userId string) (users []models.UserWithFollow, err error) {
	queryBuilder := c.db.DB.Model(&models.User{}).
		Limit(pagination.PageSize).
		Order("created_at desc").
		Select("users.*, IF((?), TRUE, FALSE) as follow_status",
			c.db.DB.Model(&models.Followers{}).
				Select("followers.created_at").
				Where(c.db.DB.Where("followers.user_id = users.id").Where("followers.follow_user_id = ?", userId)).
				Or(c.db.DB.Where("followers.follow_user_id = users.id").Where("followers.user_id = ?", userId))).
		Where("users.id != ?", userId)

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where("`users`.`email` LIKE ?", searchQuery)
	}

	if cursor != "" {
		parsedCursor, _ := time.Parse(time.RFC3339, cursor)
		queryBuilder = queryBuilder.Where("created_at < ?", parsedCursor)
	}

	return users, queryBuilder.Find(&users).Error
}

//GetOneUserById -> Get one user by id
func (c UserRepository) GetOneUserById(userId string) (user models.User, err error) {
	return user, c.db.DB.Model(&user).Where("id = ?", userId).First(&user).Error
}

func (c UserRepository) GetUsersByRoomId(roomId int64, userId string) (users []models.User, err error) {
	return users, c.db.DB.
		Model(&models.User{}).
		Joins("JOIN user_rooms UR ON users.id = UR.user_id").
		Where("user_id != ?", userId).
		Where("room_id = ?", roomId).
		Find(&users).
		Error
}
