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

func DeleteTodo (c fiber.Ctx) error {
	todoId := c.Params("id")
	userId := c.Locals("userId").(string)

	objId, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Invalid todo ID",
		})
	}

	filter := bson.M{
		"_id": objId,
		"userId": userId,
	}

	result, err := db.DB.Collection("todos").DeleteOne(c.Context(), filter) 
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete todo",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})


}

func UpdateTodo (c fiber.Ctx) error {
	type body struct {
		Title	string `json:"title"`
		Description	string `json:"description"`
		Status	string `json:"status"`
	}

	var data body
	if err := c.Bind().Body(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	update := bson.M{}

	if data.Title != ""{
		update["title"] = data.Title
	}

	if data.Description != ""{
		update["description"] = data.Description
	}

	if data.Status != ""{
		update["status"] = data.Status
	}

	if len(update) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "No fields to update",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Update sucessfully",
	})
}