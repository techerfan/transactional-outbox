package httpserver

import (
	"log"
	"order/delivery/httpserver/orderhandler"
	"order/service/orderservice"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	app          *fiber.App
	orderHandler orderhandler.Handler
}

func New(orderService *orderservice.Service) *Server {
	return &Server{
		orderHandler: orderhandler.New(orderService),
	}
}

func (s *Server) Start(address string) {
	s.app = fiber.New(fiber.Config{
		ServerHeader: "Transactional Outbox Implementation",
	})

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api := s.app.Group("/api")

	s.orderHandler.SetRoutes(api)

	if err := s.app.Listen(address); err != nil {
		log.Println("Failed to start server:", err)
	}
}
