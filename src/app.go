package src

import (
	"log"
	"main/src/db"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func SetupApp() *fiber.App {

	app := fiber.New()

	err := godotenv.Load()

	if err != nil{
		log.Fatal("Error Loading .env file")
	}

	db.ConnectDB()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}