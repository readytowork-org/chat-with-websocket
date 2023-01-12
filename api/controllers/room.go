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

type RoomController struct {
	logger          infrastructure.Logger
	roomService     services.RoomService
	userRoomService services.UserRoomService
	env             infrastructure.Env
}

func NewRoomController(logger infrastructure.Logger,
	roomService services.RoomService,
	userRoomService services.UserRoomService,
	env infrastructure.Env) RoomController {
	return RoomController{
		logger:          logger,
		roomService:     roomService,
		userRoomService: userRoomService,
		env:             env,
	}
}

func (cc RoomController) CreateRoom(c *gin.Context) {
	room := models.Room{}
	userRoom := models.UserRoom{}
	transaction := c.MustGet(constants.DBTransaction).(*gorm.DB)
	uid := c.MustGet(constants.UID).(string)

	if err := c.ShouldBindJSON(&room); err != nil {
		cc.logger.Zap.Error("Error [CreatRoom] (ShouldBindJson) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind room data")
		responses.HandleError(c, err)
		return
	}

	room, err := cc.roomService.WithTrx(transaction).CreateRoom(room)
	if err != nil {
		cc.logger.Zap.Error("Error [CreatRoom] (CreateRoom) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to Create Room")
		responses.HandleError(c, err)
		return
	}
	userRoom.UserId = uid
	userRoom.RoomId = room.ID
	err = cc.userRoomService.WithTrx(transaction).CreateUserRoom(userRoom)

	if err != nil {
		cc.logger.Zap.Error("Error [UserRoom] (userRoom) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to Create user Room")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Room Created Successfully")
}

func (cc RoomController) GetRoomWithUser(c *gin.Context) {
	ID := c.MustGet(constants.UID).(string)
	room, err := cc.roomService.GetRoomWithUser(ID)
	if err != nil {
		cc.logger.Zap.Error("Error finding users room records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users room data")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, room)

}
