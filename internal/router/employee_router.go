package router

import (
	employeesHandler "github.com/chyn-seekhachon/user-service/internal/handler/employees"
	"github.com/gofiber/fiber/v2"
)

func SetupEmployeeRoutes(api fiber.Router, employeeHandler employeesHandler.IEmployeeHandler) {
	employees := api.Group("/employee")

	employees.Post("/", employeeHandler.CreateEmployee)
	employees.Get("/", employeeHandler.GetAllEmployee)
	employees.Get("/:id", employeeHandler.GetEmployeeByID)
	employees.Put("/:id", employeeHandler.UpdateEmployee)
	employees.Delete("/:id", employeeHandler.DeleteEmployee)
}
