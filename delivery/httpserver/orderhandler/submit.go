package orderhandler

import (
	"order/dto"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) submitOrder(c *fiber.Ctx) error {
	payload := dto.SubmitOrderRequest{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	resp, err := h.orderService.Submit(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to submit order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order submitted successfully",
		"data":    resp,
	})
}
