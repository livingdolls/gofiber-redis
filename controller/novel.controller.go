package controller

import (
	"gofiber-redis/domain"
	"gofiber-redis/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NovelController struct {
	novelUseCase domain.NovelUseCase
}

func NewNovelController(novelUseCase domain.NovelUseCase) *NovelController {
	return &NovelController{novelUseCase: novelUseCase}
}

func (n *NovelController) CreateNovel(c *fiber.Ctx) error {
	var novelRequest model.Novel
	var response model.Response

	// handle request
	if err := c.BodyParser(&novelRequest); err != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	if novelRequest.Author == "" || novelRequest.Name == "" || novelRequest.Description == "" {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: "Request cannot be empty"}
		return c.Status(http.StatusBadRequest).JSON(response)
	}

	// Save database

	if err := n.novelUseCase.CreateNovel(novelRequest); err != nil {
		response = model.Response{StatusCode: http.StatusBadRequest, Message: err.Error()}
		return c.Status(http.StatusInternalServerError).JSON(response)
	}


	response = model.Response{StatusCode: http.StatusOK, Message: "Insert successfull"}
	return c.Status(http.StatusOK).JSON(response)
}

func (n *NovelController) GetNovelById(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message" : "Invalid id",
		})
	}

	novel, err := n.novelUseCase.GetNovelById(idInt)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	var res model.Response

	if novel.Name != "" {
		res = model.Response{StatusCode: http.StatusOK, Message: "Get Novel Success", Data: novel}
	} else {
		res = model.Response{StatusCode: http.StatusNotFound, Message: "Get Novel Success (null)"}

	}

	return c.Status(http.StatusOK).JSON(res)
}

func (n *NovelController) GetAllNovel(c *fiber.Ctx) error {
	var response model.Response

	novel, err := n.novelUseCase.GetAllNovel()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	response = model.Response{StatusCode: http.StatusOK, Message: "Get all novel success", Data: novel}

	return c.Status(http.StatusOK).JSON(response)

}

func (n *NovelController) DeleteNovel(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message" : "invalid id",
		})
	}

	novel, err := n.novelUseCase.DeleteNovel(idInt)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"statusCode": 200,
		"message" : "Hapus "+novel.Name+" Berhasil",
	})
}