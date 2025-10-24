package departments

import (
	"github.com/chyn-seekhachon/user-service/internal/service/departments/departmentmodel"
	"github.com/gofiber/fiber/v2"
)

func (h *DepartmentHandler) CreateDepartment(c *fiber.Ctx) error {
	var req departmentmodel.CreateDepartmentRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.departmentService.CreateDepartment(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Department created successfully",
		"data":    req,
	})
}

func (h *DepartmentHandler) GetDepartmentByID(c *fiber.Ctx) error {
	id := c.Params("id")

	department, err := h.departmentService.GetDepartmentByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Department retrieved successfully",
		"data":    department,
	})
}

func (h *DepartmentHandler) GetAllDepartment(c *fiber.Ctx) error {
	departments, err := h.departmentService.GetAllDepartment()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Departments retrieved successfully",
		"data":    departments,
	})
}

func (h *DepartmentHandler) UpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var req departmentmodel.UpdateDepartmentRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.departmentService.UpdateDepartment(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Department updated successfully",
	})
}

func (h *DepartmentHandler) DeleteDepartment(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.departmentService.DeleteDepartment(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Department deleted successfully",
	})
}
