package controllers

import (
	"main/src/models"
	"main/src/db"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	
)

func CreateTodo(c fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	
	type body struct {
		Title	string `json:"title"`
		Description	string `json:"description"`
		Status	string `json:"status"`
	}

	var data body
	if err := c.Bind().Body(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if data.Status != string(models.StatusCompleted) && data.Status != string(models.StatusIncompleted) {
		data.Status = string(models.StatusIncompleted)
	}

	todo := bson.M{
		"_id" : primitive.NewObjectID(),
		"title" : data.Title,
		"description" : data.Description,
		"status" : data.Status,
		"userId" : userId,
	}

	_, err := db.DB.Collection("todos").InsertOne(c.Context(), todo);

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cound not create Todo"})
	}

	return c.JSON(fiber.Map{
		"message": "Todo created succesfully",
		"todo": todo,
})
}

func GetTodos (c fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	cursor, err := db.DB.Collection("todos").Find(c.Context(), bson.M{
		"userId": userId,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot find todos"})
	}

	var todos []bson.M

	if err := cursor.All(c.Context(), &todos); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error" : "Cannot parse todos"})
	}

	return c.JSON(fiber.Map{"todos" : todos})
}