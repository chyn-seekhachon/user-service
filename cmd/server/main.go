package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chyn-seekhachon/user-service/internal/di"
	"github.com/chyn-seekhachon/user-service/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.leapsolutions.co.th/flow/backend/flow-library/library/envron"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "request body too large" {
				return c.Status(fiber.StatusRequestEntityTooLarge).
					JSON(fiber.Map{
						"error": "Request body size exceeds 10MB limit",
					})
			}
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"error": err.Error(),
				})
		},
	})

	if envron.GetStage() == envron.LocalStage {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	middleware.FiberMidlewareRegister(app)

	var db *gorm.DB

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       os.Getenv("DATABASE_CONNECTION_STRING"),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		errorWrapper := errors.Wrap(err, "database connection")
		log.Fatal(errorWrapper)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDb.Close()

	//inject db connection
	di.NewContainer(db, app)

	app.Get("health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	if err := app.Listen(fmt.Sprintf(":%s", envron.GetPort())); err != nil {
		log.Fatal(err)
	}
}