package middleware

import (
	"os"
)

auth jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(c fiber.Ctx) error {

	tokenString := c.cookies("jwt")

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

	if float64(time.Now().Unix()) > claims.["exp"].(float64){
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expired"})
	}
	
	c.Locals("userId", claims["userID"])

	return c.Next()

}