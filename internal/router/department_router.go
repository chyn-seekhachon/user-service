package router

import (
	departmentsHandler "github.com/chyn-seekhachon/user-service/internal/handler/departments"
	"github.com/gofiber/fiber/v2"
)

func SetupDepartmentRoutes(api fiber.Router, departmentHandler departmentsHandler.IDepartmentHandler) {
	departments := api.Group("/department")

	departments.Post("/", departmentHandler.CreateDepartment)
	departments.Get("/", departmentHandler.GetAllDepartment)
	departments.Get("/:id", departmentHandler.GetDepartmentByID)
	departments.Put("/:id", departmentHandler.UpdateDepartment)
	departments.Delete("/:id", departmentHandler.DeleteDepartment)
}
