package routes

import (
	"boilerplate-api/infrastructure"
)

type ServerRoute struct {
	wsServer *infrastructure.WsServer
	router   infrastructure.Router
}

func (i ServerRoute) Setup() {

	i.router.Gin.GET("/chat", i.wsServer.ServerWs)

}

func NewServerRoutes(
	wsServer *infrastructure.WsServer,
	router infrastructure.Router,

) ServerRoute {
	return ServerRoute{
		wsServer: wsServer,
		router:   router,
	}
}
