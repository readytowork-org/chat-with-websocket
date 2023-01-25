package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type FollowersRoutes struct {
	logger              infrastructure.Logger
	route               infrastructure.Router
	followersController controllers.FollowersController
	middleWare          middlewares.FirebaseAuthMiddleWare
	trxMiddleware       middlewares.DBTransactionMiddleware
}

func NewFollowersRoutes(
	logger infrastructure.Logger,
	route infrastructure.Router,
	followersController controllers.FollowersController,
	middleWare middlewares.FirebaseAuthMiddleWare,
	trxMiddleware middlewares.DBTransactionMiddleware,
) FollowersRoutes {
	return FollowersRoutes{
		logger:              logger,
		route:               route,
		followersController: followersController,
		middleWare:          middleWare,
		trxMiddleware:       trxMiddleware,
	}
}

func (i FollowersRoutes) Setup() {
	i.logger.Zap.Info("Setting up followers routes")
	followers := i.route.Gin.Group("/followers").Use(i.middleWare.AuthJWT())
	{
		followers.PATCH("/add/:fId", i.trxMiddleware.DBTransactionHandle(), i.followersController.AddFollower)
		followers.DELETE("/delete/:fId", i.followersController.UnFollower)
	}
}
