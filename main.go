package main

import (
	"fmt"
	"go-blog/cmd/categories"
	"go-blog/cmd/user"
	"go-blog/db"
	"go-blog/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Don't panic :D : ", r)
		}
	}()

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	app.Get("/checkAuth", func(c *fiber.Ctx) error {
		middlewares.Auth(c)
		return nil
	})
	user.Route(app, database)
	categories.Route(app, database)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	err = app.Listen(":8000")
	if err != nil {
		return
	}
}
