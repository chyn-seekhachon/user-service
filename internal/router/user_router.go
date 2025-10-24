package router

import (
	usersHandler "github.com/chyn-seekhachon/user-service/internal/handler/users"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, userHandler usersHandler.IUserHandler) {
	users := api.Group("/user")

	users.Post("/", userHandler.CreateUser)
	users.Get("/", userHandler.GetAllUser)
	users.Get("/:id", userHandler.GetUserByID)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
