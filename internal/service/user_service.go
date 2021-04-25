package service

import (
	"github.com/gofiber/fiber/v2"
	"noteme/internal/repository"
)

type User interface {
	GetUserDetail(c *fiber.Ctx) error
	UpdateUserDetail(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type user struct {
	userRepo repository.UserRepository
}

func NewUser(UserRepo repository.UserRepository) User {
	return &user{
		userRepo: UserRepo,
	}
}

func (s user) GetUserDetail(c *fiber.Ctx) error {
	user, err := s.userRepo.FindByUsername("rkritchat")
	if err != nil {
		return c.SendString("internal server error")
	}
	return c.JSON(&user)
}

func (s user) UpdateUserDetail(c *fiber.Ctx) error {
	panic("implement me")
}

func (s user) Register(c *fiber.Ctx) error {
	panic("implement me")
}

func (s user) Login(c *fiber.Ctx) error {
	panic("implement me")
}

func (s user) Logout(c *fiber.Ctx) error {
	panic("implement me")
}
