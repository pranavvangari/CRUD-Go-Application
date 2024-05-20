package main

import (
	"fmt"
	"strconv"

	"github.com/CoderVlogger/go-web-frameworks/pkg"
	"github.com/gofiber/fiber/v2"
)

type JSONTextResponse struct {
	Message string
}

var (
	entitiesRepo pkg.EntityRepository = pkg.NewEntityMemoryRepository()
	pageSize     int                  = 4
)

func main() {
	fmt.Println("hello world!")

	entitiesRepo.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error { // c for context
		return c.JSON(JSONTextResponse{Message: "Hello, Fiber!"})
	})

	entitiesAPI := app.Group("/entities")

	entitiesAPI.Get("/", entitiesList)
	entitiesAPI.Get("/:id", entitiesGet)
	entitiesAPI.Post("/", entitiesAdd)
	entitiesAPI.Put("/", entitiesUpdate)
	entitiesAPI.Delete("/:id", entitiesDelete)

	app.Listen(":8080")

}

func entitiesList(c *fiber.Ctx) error { // when a page number is entered into the url as a get request

	pageStr := c.Query("page", "1")    //pageStr will store either the page number entered, or 1 if no page number is entered
	page, err := strconv.Atoi(pageStr) //strconv is short for string conversion; strconv.Atoi is like parseInteger, where it changes a stirng into an int and stores it in page
	if err != nil {
		page = 1
	}
	entities, err := entitiesRepo.List(page, pageSize) //start from the specified page number
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()} //using a struct in our model from pkg
		return c.Status(fiber.StatusNotFound).JSON(errMsg)
	}
	return c.JSON(entities)
}

func entitiesGet(c *fiber.Ctx) error {
	entityID := c.Params("id")

	entity, err := entitiesRepo.Get(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()} //using a struct in our model from pkg
		return c.Status(fiber.StatusNotFound).JSON(errMsg)
	}
	return c.JSON(entity)

}

func entitiesAdd(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity) //putting that json POST method information into entity
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()} //using a struct in our model from pkg
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = entitiesRepo.Add(&entity)

	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()} //using a struct in our model from pkg
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	return c.JSON(entity)
}

func entitiesUpdate(c *fiber.Ctx) error {
	var entity pkg.Entity

	err := c.BodyParser(&entity) //putting that json POST method information into entity
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()} //using a struct in our model from pkg
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = entitiesRepo.Update(&entity)

	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()} //using a struct in our model from pkg
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	return c.JSON(entity)
}

func entitiesDelete(c *fiber.Ctx) error {
	entityID := c.Params("id")

	err := entitiesRepo.Delete(entityID)
	if err != nil {
		errMsg := pkg.TextResponse{Message: err.Error()} //using a struct in our model from pkg
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}
	return c.JSON(pkg.TextResponse{Message: "entity deleted"})

}
