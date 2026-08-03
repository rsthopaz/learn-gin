package main

import (
	"main/src"
	"log"
)

func main() {
	app := src.SetupApp()

	port := ":3000"
	log.Println("Server is running on port " + port)
	app.Listen(port)
}