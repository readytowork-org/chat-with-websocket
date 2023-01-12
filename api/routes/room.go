package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type RoomRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	roomController controllers.RoomController
	middleWare     middlewares.FirebaseAuthMiddleWare
	trxMiddleware  middlewares.DBTransactionMiddleware
}

func (i RoomRoutes) Setup() {
	i.logger.Zap.Info("Setting up room routes")
	rooms := i.router.Gin.Group("/rooms").Use(i.middleWare.AuthJWT())
	{
		rooms.POST("/create", i.trxMiddleware.DBTransactionHandle(), i.roomController.CreateRoom)
		rooms.GET("/getRooms", i.roomController.GetRoomWithUser)
	}
}

func NewRoomRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	roomController controllers.RoomController,
	middleWare middlewares.FirebaseAuthMiddleWare,
	trxMiddleware middlewares.DBTransactionMiddleware,
) RoomRoutes {
	return RoomRoutes{
		logger:         logger,
		router:         router,
		roomController: roomController,
		trxMiddleware:  trxMiddleware,
		middleWare:     middleWare,
	}
}
