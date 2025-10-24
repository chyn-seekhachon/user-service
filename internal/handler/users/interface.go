package users

import "github.com/gofiber/fiber/v2"

type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}
