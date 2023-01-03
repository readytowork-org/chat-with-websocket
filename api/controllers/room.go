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
	logger      infrastructure.Logger
	roomService services.RoomService
	env         infrastructure.Env
}

func NewRoomController(logger infrastructure.Logger,
	roomService services.RoomService,
	env infrastructure.Env) RoomController {
	return RoomController{
		logger:      logger,
		roomService: roomService,
		env:         env,
	}
}

func (cc RoomController) CreateRoom(c *gin.Context) {
	room := models.Room{}

	if err := c.ShouldBindJSON(&room); err != nil {
		cc.logger.Zap.Error("Error [CreatRoom] (ShouldBindJson) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind room data")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Room Created Successfully")
}
