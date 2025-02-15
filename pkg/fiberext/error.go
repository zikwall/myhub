package fiberext

import (
	"github.com/gofiber/fiber/v2"

	"github.com/zikwall/myhub/pkg/log"
)

const defaultErrorMessage = "Internal Server Error"

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	if err != nil {
		code := fiber.StatusInternalServerError
		value := defaultErrorMessage

		log.Warningf("error handler: %s", err)

		return ctx.Status(code).JSON(fiber.Map{
			"status":  code,
			"message": value,
		})
	}

	return nil
}
