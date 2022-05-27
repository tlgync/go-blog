package user

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Get(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Login(*fiber.Ctx) error
}

type handler struct {
	service Service
}

var _ Handler = handler{}

func NewHandler(service Service) Handler {
	return handler{service: service}

}

type Response struct {
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (h handler) Login(c *fiber.Ctx) error {

	data := LoginDTO{}

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	_, err = h.service.Login(data)

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(Response{Data: tokenString, Status: true, Message: "Login is successfully"})
}

func (h handler) Get(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	model, err := h.service.Get(uint(id))

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	return c.Status(200).JSON(Response{Data: model, Status: true})
}

func (h handler) Create(c *fiber.Ctx) error {
	model := Model{}
	err := c.BodyParser(&model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	_, err = h.service.Create(model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	return c.Status(201).JSON(Response{Data: model, Status: true, Message: "User created"})
}

func (h handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	_, err = h.service.Delete(uint(id))

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	return c.Status(200).JSON(Response{Status: true, Message: "User deleted"})
}
