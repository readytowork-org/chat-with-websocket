package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/constants"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (cc RoomController) GetRoomWithUser(c *gin.Context) {
	ID := c.MustGet(constants.UID).(string)
	cursor := c.Query("cursor")
	room, err := cc.roomService.GetRoomWithUser(ID, cursor)
	if err != nil {
		cc.logger.Zap.Error("Error finding users room records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users room data")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, room)
}

func (cc RoomController) GetRoomsMessages(c *gin.Context) {
	roomId, _ := strconv.ParseInt(c.Param("room-id"), 10, 64)
	cursor := c.Query("cursor")
	messages, err := cc.messageService.GetMessageWithUser(roomId, cursor)
	if err != nil {
		cc.logger.Zap.Error("Error finding user room's message", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users room message")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, messages)
}
