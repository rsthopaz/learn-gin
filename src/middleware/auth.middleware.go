package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v3"

)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(c fiber.Ctx) error {

	tokenString := c.Cookies("jwt")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims := token.Claims.(jwt.MapClaims)

	if float64(time.Now().Unix()) > claims["exp"].(float64){
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expired"})
	}
	
	c.Locals("userId", claims["userId"])

	return c.Next()

}