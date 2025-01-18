package routers

import (
	"S3/handlers"

	"github.com/gofiber/fiber/v2"
)

func S3Routes(app *fiber.App) {
	api := app.Group("/S3")

	api.Post("/", handlers.UploadFile)
	api.Delete("/:fileName", handlers.DeleteFile)
}
