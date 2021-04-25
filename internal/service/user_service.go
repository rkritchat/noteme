package service

import (
	"github.com/gofiber/fiber/v2"
	"noteme/internal/repository"
)

type User interface {
	GetUser(c *fiber.Ctx) error
}

type user struct {
	userRepo repository.UserRepository
}

func NewUser(UserRepo repository.UserRepository) User {
	return &user{
		userRepo: UserRepo,
	}
}

func (s user) GetUser(c *fiber.Ctx) error {
	user, err := s.userRepo.FindByUsername("rkritchat")
	if err != nil {
		return c.SendString("internal server error")
	}
	return c.JSON(&user)
}
