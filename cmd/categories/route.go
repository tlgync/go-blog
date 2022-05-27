package categories

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Route(app *fiber.App, database *gorm.DB) {
	repoCategory := NewReporsitory(database)
	categoryService := NewService(repoCategory)
	handler := NewHandler(categoryService)

	category := app.Group("/category")
	category.Get("/getAllCategory", handler.GetAll)
	category.Get("/category/:id", handler.Get)
	category.Post("/category/", handler.Create)
	category.Get("/category/delete/:id", handler.Delete)
}
