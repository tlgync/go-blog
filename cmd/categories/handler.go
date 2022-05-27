package categories

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
}

type Handler interface {
	Get(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	GetAll(*fiber.Ctx) error
}

type handler struct {
	service Service
}

var _ Handler = handler{}

func NewHandler(service Service) Handler {
	return handler{service: service}
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

	return c.Status(200).JSON(Response{Data: model, Status: false})
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

	return c.Status(201).JSON(Response{Data: model, Status: true, Message: "Category created"})
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

	return c.Status(200).JSON(Response{Status: true, Message: "Category deleted"})
}

func (h handler) GetAll(c *fiber.Ctx) error {
	models, err := h.service.GetAll()

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error(), Status: false})
	}

	return c.Status(200).JSON(Response{Data: models, Status: true})
}
