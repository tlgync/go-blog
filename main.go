package main

import (
	"go-blog/cmd/user"
	"go-blog/cmd/user/middlewares"
	"go-blog/db"
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
	err = repo.Migration()
	if err != nil {
		log.Fatal(err)
	}
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()

	app.Get("/users/:id", handler.Get)
	app.Post("/users/", handler.Create)
	app.Get("/users/delete/:id", handler.Delete)
	app.Get("/checkAuth", func(c *fiber.Ctx) error {
		middlewares.Auth(c)
		return nil
	})
	app.Post("/users/login", handler.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Listen(":8000")
}
