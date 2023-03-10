package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FollowersController struct {
	logger           infrastructure.Logger
	followersService services.FollowersService
	roomService      services.RoomService
	userRoomService  services.UserRoomService
	env              infrastructure.Env
}

func NewFollowersController(
	logger infrastructure.Logger,
	followersService services.FollowersService,
	roomService services.RoomService,
	userRoomService services.UserRoomService,
	env infrastructure.Env,

) FollowersController {
	return FollowersController{
		logger:           logger,
		followersService: followersService,
		roomService:      roomService,
		userRoomService:  userRoomService,
		env:              env,
	}
}

func (cc FollowersController) AddFollower(c *gin.Context) {
	transaction := c.MustGet(constants.DBTransaction).(*gorm.DB)

	followers := models.Followers{
		UserId:       c.MustGet(constants.UID).(string),
		FollowUserId: c.Param("fId"),
	}

	followers, err := cc.followersService.WithTrx(transaction).AddFollower(followers)
	if err != nil {
		cc.logger.Zap.Error("Error [AddFollower] (AddFollowers) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to add friend")
		responses.HandleError(c, err)
		return
	}

	userRoom, err := cc.userRoomService.WithTrx(transaction).GetUserRoomByFollowId(followers.ID)
	if err != nil {
		room := models.Room{Name: "", IsPrivate: true}
		room, err = cc.roomService.WithTrx(transaction).CreateRoom(room)
		if err != nil {
			cc.logger.Zap.Error("Error [CreatRoom] (CreateRoom) :", err)
			err := errors.BadRequest.Wrap(err, "Failed to Create Room")
			responses.HandleError(c, err)
			return
		}

		userRoom = models.UserRoom{FollowerId: followers.ID, RoomId: room.ID}
		err = cc.userRoomService.WithTrx(transaction).CreateUserRoom(userRoom)
		if err != nil {
			cc.logger.Zap.Error("Error [UserRoom] (userRoom) :", err)
			err := errors.BadRequest.Wrap(err, "Failed to Create user Room")
			responses.HandleError(c, err)
			return
		}
	}

	responses.SuccessJSON(c, http.StatusOK, "Friend Added & Room Created Successfully")
}

func (cc FollowersController) UnFollower(c *gin.Context) {
	followers := models.Followers{
		UserId:       c.MustGet(constants.UID).(string),
		FollowUserId: c.Param("fId"),
	}

	followers, err := cc.followersService.UnFollower(followers)
	if err != nil {
		cc.logger.Zap.Error("Error [UnFollow Friend] (UnFollow Followers) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to UnFollow friend")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Friend Removed")
}
