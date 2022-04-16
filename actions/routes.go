package actions

import (
	authuser "livegift_back/actions/auth"
	"livegift_back/actions/events"
	"livegift_back/actions/gifts"
	"livegift_back/actions/middleware/authorization"
	"livegift_back/actions/users"

	"github.com/gobuffalo/buffalo"
)

func SetRoutes(app *buffalo.App) *buffalo.App {
	app.GET("/", HomeHandler)

	authRoutes := app.Group("/auth")
	authRoutes.POST("/login", authuser.AuthLogin)
	authRoutes.POST("/signup", authuser.AuthRegister)
	authRoutes.Middleware.Remove(authorization.Authorizator)

	giftRoutes := app.Group("/gifts")
	giftRoutes.GET("/", gifts.ListGift)
	giftRoutes.GET("/qr/view", gifts.GenerateQRCode)
	giftRoutes.POST("/create", gifts.CreateGift)

	userRoutes := app.Group("/users")
	userRoutes.GET("/", users.Index)

	eventRoutes := app.Group("/events")
	eventRoutes.GET("/", events.Index)
	eventRoutes.POST("/create", events.Create)

	return app
}
