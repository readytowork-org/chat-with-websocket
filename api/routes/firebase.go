package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type FirebaseRoutes struct {
	firebaseController controllers.FirebaseController
	router             infrastructure.Router
	firebaseAuth       middlewares.FirebaseAuthMiddleWare
}

func NewFirebaseRoutes(
	firebaseController controllers.FirebaseController,
	router infrastructure.Router) FirebaseRoutes {
	return FirebaseRoutes{
		firebaseController: firebaseController,
		router:             router,
	}
}

func (fr FirebaseRoutes) Setup() {
	firebase := fr.router.Gin.Group("fb").Use(fr.firebaseAuth.AuthJWT())
	{
		firebase.POST("/create", fr.firebaseController.CreateUser)
	}
}
