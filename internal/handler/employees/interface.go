package employees

import "github.com/gofiber/fiber/v2"

type IEmployeeHandler interface {
	CreateEmployee(c *fiber.Ctx) error
	GetEmployeeByID(c *fiber.Ctx) error
	GetAllEmployee(c *fiber.Ctx) error
	UpdateEmployee(c *fiber.Ctx) error
	DeleteEmployee(c *fiber.Ctx) error
}
