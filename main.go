package main

import (
	"go-blog/cmd/categories"
	"go-blog/cmd/user"
	"go-blog/db"
	"go-blog/middlewares"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type JWTClaimm struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func main() {

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := user.NewReporsitory(database)
	repoCategory := categories.NewReporsitory(database)

	service := user.NewService(repo)
	handler := user.NewHandler(service)

	categoryService := categories.NewService(repoCategory)
	handlerCategory := categories.NewHandler(categoryService)

	app := fiber.New()

	app.Get("/users/:id", handler.Get)
	app.Post("/users/", handler.Create)
	app.Get("/users/delete/:id", handler.Delete)
	app.Post("/users/login", handler.Login)
	app.Get("/checkAuth", func(c *fiber.Ctx) error {
		middlewares.Auth(c)
		return nil
	})

	app.Get("/category/:id", handlerCategory.Get)
	app.Post("/category/", handlerCategory.Create)
	app.Get("/category/delete/:id", handlerCategory.Delete)
	app.Get("/getAllCategory", handlerCategory.GetAll)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Listen(":8000")
}
