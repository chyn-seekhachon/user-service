package di

import (
	"context"
	"time"

	"gorm.io/gorm"

	// Repositories
	departmentsRepo "github.com/chyn-seekhachon/user-service/internal/repository/departments"
	employeesRepo "github.com/chyn-seekhachon/user-service/internal/repository/employees"
	usersRepo "github.com/chyn-seekhachon/user-service/internal/repository/users"
	"github.com/chyn-seekhachon/user-service/internal/router"
	"github.com/gofiber/fiber/v2"

	// Services
	departmentsService "github.com/chyn-seekhachon/user-service/internal/service/departments"
	employeesService "github.com/chyn-seekhachon/user-service/internal/service/employees"
	usersService "github.com/chyn-seekhachon/user-service/internal/service/users"

	// Handlers
	departmentsHandler "github.com/chyn-seekhachon/user-service/internal/handler/departments"
	employeesHandler "github.com/chyn-seekhachon/user-service/internal/handler/employees"
	usersHandler "github.com/chyn-seekhachon/user-service/internal/handler/users"
)

type Container struct {
	// Database
	DB *gorm.DB

	// Repositories
	UserRepository       usersRepo.IUserRepository
	DepartmentRepository departmentsRepo.IDepartmentRepository
	EmployeeRepository   employeesRepo.IEmployeeRepository

	// Services
	UserService       usersService.IUserService
	DepartmentService departmentsService.IDepartmentService
	EmployeeService   employeesService.IEmployeeService

	// Handlers
	UserHandler       usersHandler.IUserHandler
	DepartmentHandler departmentsHandler.IDepartmentHandler
	EmployeeHandler   employeesHandler.IEmployeeHandler
}

// NewContainer creates and wires all dependencies
func NewContainer(db *gorm.DB, app *fiber.App) *Container {
	c := &Container{
		DB: db,
	}

	// Dependency injection flow: DB → Repository → Service → Handler
	c.SetUpRepository(db)
	c.SetUpService()
	c.SetUpHandler()
	c.SetupRoutes(app)
	return c
}

func (c *Container) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (c *Container) SetupRoutes(app *fiber.App) {
	// Health check endpoint
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status":  "ok",
			"service": "user-service",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Setup domain routes by injecting router and handlers
	router.SetupUserRoutes(api, c.UserHandler)
	router.SetupDepartmentRoutes(api, c.DepartmentHandler)
	router.SetupEmployeeRoutes(api, c.EmployeeHandler)
}

func (c *Container) SetUpRepository(db *gorm.DB) {
	c.UserRepository = usersRepo.NewUserRepository(db, c.withTimeout)
	c.DepartmentRepository = departmentsRepo.NewDepartmentRepository(db, c.withTimeout)
	c.EmployeeRepository = employeesRepo.NewEmployeeRepository(db, c.withTimeout)
}

func (c *Container) SetUpService() {
	c.UserService = usersService.NewUserService(c.UserRepository)
	c.DepartmentService = departmentsService.NewDepartmentService(c.DepartmentRepository)
	c.EmployeeService = employeesService.NewEmployeeService(c.EmployeeRepository)
}

func (c *Container) SetUpHandler() {
	c.UserHandler = usersHandler.NewUserHandler(c.UserService)
	c.DepartmentHandler = departmentsHandler.NewDepartmentHandler(c.DepartmentService)
	c.EmployeeHandler = employeesHandler.NewEmployeeHandler(c.EmployeeService)
}

