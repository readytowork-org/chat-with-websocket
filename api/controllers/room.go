package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if err := c.ShouldBindJSON(&room); err != nil {
		cc.logger.Zap.Error("Error [CreatRoom] (ShouldBindJson) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind room data")
		responses.HandleError(c, err)
		return
	}

	err := cc.roomService.CreateRoom(room)
	if err != nil {
		cc.logger.Zap.Error("Error [CreatRoom] (CreateRoom) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to Create Room")
		responses.HandleError(c, err)
		return
	}
	cc.userRoomService.CreateUserRoom(userRoom)

	responses.SuccessJSON(c, http.StatusOK, "Room Created Successfully")
}
