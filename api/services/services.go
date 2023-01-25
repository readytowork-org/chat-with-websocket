package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewRoomService),
	fx.Provide(NewUserRoomService),
	fx.Provide(NewFirebaseService),
	fx.Provide(NewMessageService),
	fx.Provide(NewFollowersService),
)
