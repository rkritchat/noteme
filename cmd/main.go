package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"time"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/test", func(c *fiber.Ctx) error {
		time.Sleep(1 * time.Second)

		return c.SendString("Hello world")
	})
	log.Fatal(app.Listen(":3000"))
}
