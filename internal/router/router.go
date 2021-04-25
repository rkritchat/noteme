package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"noteme/internal/service"
)

func NewRouter(userService service.User) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/test", userService.GetUser)
	return app
}
