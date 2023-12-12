package resources

import (
	"github.com/gofiber/fiber/v2"
	"neema.co.za/rest/utils/models"
)

func SendResponse[T any](objet []T, message string, c *fiber.Ctx, op string) {
	basemessage := models.Basemessage[T]{
		Success: true,
		Data:    objet,
		Message: message,
	}

	switch op {
	case "create":
		c.Status(fiber.StatusCreated).JSON(fiber.Map{"response": basemessage})
		return
	default:
		c.Status(fiber.StatusOK).JSON(fiber.Map{"response": basemessage})
		return
	}

}

func SendResponses[T any](objet T, message string, c *fiber.Ctx, op string) {
	basemessage := models.Basemessageone[T]{
		Success: true,
		Data:    objet,
		Message: message,
	}

	switch op {
	case "create":
		c.Status(fiber.StatusCreated).JSON(fiber.Map{"response": basemessage})
		return
	default:
		c.Status(fiber.StatusOK).JSON(fiber.Map{"response": basemessage})
		return
	}

}

func SendError(err error, c *fiber.Ctx) {
	basemessagerror := models.Basemessageerror{
		Success: false,
		Error:   err.Error(),
	}

	c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"response": basemessagerror})
}

func SendErrors(err string, c *fiber.Ctx) {
	basemessagerror := models.Basemessageerror{
		Success: false,
		Error:   err,
	}

	c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"response": basemessagerror})
}