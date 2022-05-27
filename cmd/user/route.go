package user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, database *gorm.DB) {
	repo := NewReporsitory(database)
	service := NewService(repo)
	handler := NewHandler(service)

	app.Get("/users/:id", handler.Get)
	app.Post("/users/", handler.Create)
	app.Get("/users/delete/:id", handler.Delete)
	app.Post("/users/login", handler.Login)
}
