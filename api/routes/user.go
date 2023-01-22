package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

// UserRoutes -> struct
type UserRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	userController controllers.UserController
	trxMiddleware  middlewares.DBTransactionMiddleware
	firebaseAuth   middlewares.FirebaseAuthMiddleWare
}

// Setup user routes
func (i UserRoutes) Setup() {
	i.logger.Zap.Info(" Setting up user routes")

	i.router.Gin.POST("/create", i.trxMiddleware.DBTransactionHandle(), i.userController.CreateUser)

	users := i.router.Gin.Group("/users").Use(i.firebaseAuth.AuthJWT())
	{
		users.GET("/get-all/:cursor", i.userController.GetAllUsers)
	}
}

// NewUserRoutes -> creates new user controller
func NewUserRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	userController controllers.UserController,
	trxMiddleware middlewares.DBTransactionMiddleware,
	firebaseAuth middlewares.FirebaseAuthMiddleWare,
) UserRoutes {
	return UserRoutes{
		router:         router,
		logger:         logger,
		userController: userController,
		trxMiddleware:  trxMiddleware,
		firebaseAuth:   firebaseAuth,
	}
}
