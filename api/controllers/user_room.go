package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"

	"github.com/gin-gonic/gin"
)

type UseRoomController struct {
	logger          infrastructure.Logger
	userRoomService services.UserRoomService
	env             infrastructure.Env
}

func NewUserRoomController(logger infrastructure.Logger,
	userRoomService services.UserRoomService,
	env infrastructure.Env) UseRoomController {
	return UseRoomController{
		logger:          logger,
		userRoomService: userRoomService,
		env:             env,
	}
}

func (cc UseRoomController) CreateUserRoom(c *gin.Context) {
	userRoom := models.UserRoom{}

	if err := c.ShouldBindJSON(&userRoom); err != nil {
		cc.logger.Zap.Error("Error [CreatUserRoom] (ShouldBindJson) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind userRoom data")
		responses.HandleError(c, err)
		return
	}

	err := cc.userRoomService.CreateUserRoom(userRoom)
	if err != nil {
		cc.logger.Zap.Error("Error [CreatUserRoom] (CreateUserRoom) :", err)
		err := errors.BadRequest.Wrap(err, "Failed to User Create Room")
		responses.HandleError(c, err)
		return
	}

}
