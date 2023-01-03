package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/infrastructure"
)

type RoomRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	roomController controllers.RoomController
}

func (i RoomRoutes) Setup() {
	i.logger.Zap.Info("Setting up room routes")
	rooms := i.router.Gin.Group("/rooms")
	{
		rooms.POST("", i.roomController.CreateRoom)
	}
}

func NewRoomRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	roomController controllers.RoomController,
) RoomRoutes {
	return RoomRoutes{
		logger:         logger,
		router:         router,
		roomController: roomController,
	}
}
