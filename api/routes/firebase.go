package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/infrastructure"
)

type FirebaseRoutes struct {
	firebaseController controllers.FirebaseController
	router             infrastructure.Router
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
	firebase := fr.router.Gin.Group("fb").Use()
	{
		firebase.GET("/", fr.firebaseController.CreateUser)
	}
}
