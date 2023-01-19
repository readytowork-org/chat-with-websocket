package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomController struct {
	logger          infrastructure.Logger
	roomService     services.RoomService
	userRoomService services.UserRoomService
	messageService  services.MessageService
	env             infrastructure.Env
}

func NewRoomController(logger infrastructure.Logger,
	roomService services.RoomService,
	userRoomService services.UserRoomService,
	messageService services.MessageService,
	env infrastructure.Env) RoomController {
	return RoomController{
		logger:          logger,
		roomService:     roomService,
		userRoomService: userRoomService,
		env:             env,
		messageService:  messageService,
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
	cursor := c.Param("cursor")
	room, err := cc.roomService.GetRoomWithUser(ID, cursor)
	if err != nil {
		cc.logger.Zap.Error("Error finding users room records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users room data")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, room)

}

func (cc RoomController) CreateMessageWithUser(c *gin.Context) {
	message := models.Message{}
	uid := c.MustGet(constants.UID).(string)
	roomId, _ := strconv.ParseInt(c.Param("room-id"), 10, 64)

	if err := c.ShouldBindJSON(&message); err != nil {
		cc.logger.Zap.Error("Error [CreatMessage] (ShouldBindJson) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind message data")
		responses.HandleError(c, err)
		return
	}

	message.RoomId = roomId
	message.UserId = uid
	err := cc.messageService.CreateMessageWithUser(roomId, message)

	if err != nil {
		cc.logger.Zap.Error("Error [CreatMessage] (CreateMessage) :", err.Error())
		err := errors.BadRequest.Wrap(err, "Failed to Create Message")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, message)

}

func (cc RoomController) GetRoomsMessages(c *gin.Context) {
	cursor := c.Param("cursor")

	roomId, _ := strconv.ParseInt(c.Param("room-id"), 10, 64)

	messages, err := cc.messageService.GetMessageWithUser(roomId, cursor)
	if err != nil {
		cc.logger.Zap.Error("Error finding user room's message", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users room message")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, messages)

}
