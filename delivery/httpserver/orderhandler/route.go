package orderhandler

import "github.com/gofiber/fiber/v2"

func (h Handler) SetRoutes(app fiber.Router) {
	app.Post("/orders", h.submitOrder)
}
