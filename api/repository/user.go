package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"time"

	"gorm.io/gorm"
)

// UserRepository -> database structure
type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository -> creates a new User repository
func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx enables repository with transaction
func (c UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transaction Database not found in gin context. ")
		return c
	}
	c.db.DB = trxHandle
	return c
}

// Save -> User
func (c UserRepository) Create(User models.User) (models.User, error) {
	return User, c.db.DB.Create(&User).Error
}

// GetAllUser -> Get All users
func (c UserRepository) GetAllUsers(pagination utils.Pagination, cursor string, userId string) (users []models.UserWithFalano, totalRows int64, err error) {
	totalRows = 0
	time, _ := time.Parse(time.RFC3339, cursor)

	followers := c.db.DB

	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.User{}).
		Select("users.* , (?) as follow_status", followers.Model(&models.Followers{}).
			Select("IF (followers.user_id IS NOT NULL, true, false)").
			Where("followers.user_id = ?", userId).Or("followers.follow_user_id = ?", userId).Limit(1)).
		Where(" users.id != ?", userId)

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`users`.`email` LIKE ?", searchQuery))
	}

	if cursor != "" {
		queryBuilder = queryBuilder.Where("created_at < ?", time)
	}

	err = queryBuilder.
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return users, totalRows, err
}
