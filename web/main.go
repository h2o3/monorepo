package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Status string `json:"status"`
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	root := app.Group("/")

	root.Get("/", func(c *fiber.Ctx) error {
		if c.Query("name") == "" {
			return fmt.Errorf("invalid argument")
		}

		return c.JSON(Response{
			Status: "OK",
		})
	})

	log.Fatal(app.Listen(":9000"))
}
