package routes

import (

	"main/src/controllers"
	"main/src/middleware"
	"github.com/gofiber/fiber/v3"
)

func TodoRoutes(app *fiber.App) {
	auth := app.Group("/todo", middleware.AuthMiddleware)
	auth.Post("/", controllers.CreateTodo)
	auth.Get("/", controllers.GetTodos)
	auth.Delete("/:id", controllers.DeleteTodo)
	auth.Put("/:id", controllers.UpdateTodo)
	
}