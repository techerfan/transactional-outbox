package orderhandler

import "order/service/orderservice"

type Handler struct {
	orderService *orderservice.Service
}

func New(orderService *orderservice.Service) Handler {
	return Handler{
		orderService: orderService,
	}
}
