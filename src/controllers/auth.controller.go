package controllers

import (
	"main/src/models"
	"os"
	"time"
	"main/src/db"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func RegisterUser(c fiber.Ctx) error {
	type request struct {
		Email   string `json:"email"`
		Password string `json:"password"`
	}


var body request

if err := c.BodyParser(&body); err != nil {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
}

var existingUser models.User

err := db.DB.Collection("users").FindOne(c.Context(), bson.M{"email": body.Email}).Decode(&existingUser)
if err == nil {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
}

hashPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

userDoc := bson.M{
	"email": body.Email,
	"password": string(hashPassword),
}

_, err = db.DB.Collection("users").InsertOne(c.Context(), userDoc)
if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
}

return c.JSON(fiber.Map{"message": "User registered successfully"})

}

func LoginUser(c fiber.Ctx) error {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": " Cannot parse request"})
	}

	var user models.User

	err := db.DB.Collection("users").FindOne(c.Context(), bson.M{"email": body.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password));
	err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaimns{
		"userId": user.ID.Hex(),
		"exp": time.Now().Add(24 * time.hour).Unix(),
	})

	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": " Could not login"})
	}

	c.Cookie(&fiber.Cookie{
		Name: "jwt",
		Value: t,
		Expires: time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "Logged in succesfuly"})
}