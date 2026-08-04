package routes

import (

	"main/src/controllers"
	"main/src/middleware"
	"github.com/gofiber/fiber/v3"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)
	auth.Post("/logout", middleware.AuthMiddleware, controllers.LogoutUser)
}