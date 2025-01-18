package main

import (
	"S3/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())

	routers.S3Routes(app)

	log.Fatal(app.Listen(":8000"))
}
