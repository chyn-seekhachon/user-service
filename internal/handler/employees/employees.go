package employees

import (
	"github.com/chyn-seekhachon/user-service/internal/service/employees/employeemodel"
	"github.com/gofiber/fiber/v2"
)

func (h *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var req employeemodel.CreateEmployeeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.employeeService.CreateEmployee(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Employee created successfully",
		"data":    req,
	})
}

func (h *EmployeeHandler) GetEmployeeByID(c *fiber.Ctx) error {
	id := c.Params("id")

	employee, err := h.employeeService.GetEmployeeByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employee retrieved successfully",
		"data":    employee,
	})
}

func (h *EmployeeHandler) GetAllEmployee(c *fiber.Ctx) error {
	employees, err := h.employeeService.GetAllEmployee()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employees retrieved successfully",
		"data":    employees,
	})
}

func (h *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var req employeemodel.UpdateEmployeeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.employeeService.UpdateEmployee(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employee updated successfully",
	})
}

func (h *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.employeeService.DeleteEmployee(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Employee deleted successfully",
	})
}
