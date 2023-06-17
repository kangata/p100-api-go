package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/kangata/p100-api-go/controllers"
)

var port string

func main() {
	loadEnv()

	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("APP_EMAIL"): os.Getenv("APP_PASSWORD"),
		},
		Unauthorized: func(c *fiber.Ctx) error {
			c.Status(fiber.ErrUnauthorized.Code)

			return c.JSON(fiber.Map{
				"code":    fiber.ErrUnauthorized.Code,
				"message": fiber.ErrUnauthorized.Message,
			})
		},
	}))

	registerRoutes(app)

	app.Listen(fmt.Sprintf(":%s", port))
}

func loadEnv() {
	port = os.Getenv("APP_PORT")

	if port == "" {
		port = "3000"
	}
}

func registerRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": fiber.StatusOK, "message": "Online"})
	})

	app.Get("/device", controllers.GetDevice)
	app.Post("/switch", controllers.SwitchDeviceStatus)
}
