package myhub

import "github.com/gofiber/fiber/v2"

type UseCase interface {
}

type Handler struct {
	useCase UseCase
}

func New(useCase UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) MountRoutes(app *fiber.App) {
	app.Get("/alive", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Alive")
	})
}
