package router

import (
	"gofiber-redis/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, novelController *controller.NovelController) *fiber.App {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	router.Post("/novel", novelController.CreateNovel)
	router.Get("/novel/:id", novelController.GetNovelById)
	router.Get("/novels", novelController.GetAllNovel)
	router.Delete("/novel/:id", novelController.DeleteNovel)
	return router;
}