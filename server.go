package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./comic-reader-front/build/")

	app.Listen(":4040")
}
