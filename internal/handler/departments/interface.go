package departments

import "github.com/gofiber/fiber/v2"

type IDepartmentHandler interface {
	CreateDepartment(c *fiber.Ctx) error
	GetDepartmentByID(c *fiber.Ctx) error
	GetAllDepartment(c *fiber.Ctx) error
	UpdateDepartment(c *fiber.Ctx) error
	DeleteDepartment(c *fiber.Ctx) error
}
