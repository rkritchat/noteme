package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"noteme/internal/service"
)

func NewRouter(userService service.User) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Mount("/user", userRouter(userService))
	return app
}

func userRouter(userService service.User) *fiber.App {
	u := fiber.New()
	u.Get("/detail", userService.GetUserDetail)
	u.Patch("/detail", userService.UpdateUserDetail)
	u.Post("/register", userService.Register)
	u.Post("/login", userService.Login)
	u.Post("/logout", userService.Logout)
	return u
}
