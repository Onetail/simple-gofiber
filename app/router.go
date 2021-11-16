package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func setRouters(router *fiber.App) {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

}

func listenPort(router *fiber.App, port int) {
	appPort := fmt.Sprintf(":%d", port)
	router.Listen(appPort)
}

func StartServer(router *fiber.App, port int) {
	setRouters(router)
	listenPort(router, port)
}
