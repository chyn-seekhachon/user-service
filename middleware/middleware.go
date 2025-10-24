package middleware

import (
	"strings"

	"gitlab.leapsolutions.co.th/flow/backend/flow-library/library/envron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMidlewareRegister(app *fiber.App) {
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	if envron.GetStage() == envron.ProductionStage || envron.GetStage() == envron.LocalStage {
		app.Use(logger.New(logger.Config{
			Format:     "${time} ${ip} ${status} ${method} ${path} ${latency}\n",
			TimeFormat: "02-Jan-2006 15:04:05 ",
			TimeZone:   "Asia/Bangkok",
		}))
	}

	app.Use(compress.New())
	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			return strings.Contains(c.Path(), "vansales") || strings.Contains(c.Path(), "pos")
		},
		Level: compress.LevelBestSpeed,
	}))
}
