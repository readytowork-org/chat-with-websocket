package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoomRoutes),
	fx.Provide(NewServerRoutes),
	fx.Provide(NewFirebaseRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoutes,
	roomRoutes RoomRoutes,
	serverRoutes ServerRoute,
	firebaseRoutes FirebaseRoutes,
) Routes {
	return Routes{
		userRoutes,
		roomRoutes,
		serverRoutes,
		firebaseRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
