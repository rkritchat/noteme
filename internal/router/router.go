package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
	"noteme/internal/service"
)

func NewRouter(userService service.User) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	group := app.Group("/user")
	group.Add(http.MethodGet, "/test", userService.GetUser)
	return app
}
