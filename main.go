package main

import (
	"simple-gofiber/app"

	"github.com/gofiber/fiber/v2"
)

func main() {
	application := fiber.New()
	app.StartServer(application, 3000)

}
