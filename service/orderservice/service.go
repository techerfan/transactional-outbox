package orderservice

import (
	"context"
	"order/contract/broker"
	"order/entity"
)

type Repository interface {
	GetOrdersFromOutbox(ctx context.Context, limit int) ([]entity.OrderOutbox, error)
	UpdateOrderOutbox(ctx context.Context, orderOutbox entity.OrderOutbox) error
	CreateOrder(ctx context.Context, orderEntity entity.Order) (entity.Order, entity.OrderOutbox, error)
	FindOrderByID(ctx context.Context, id uint) (entity.Order, error)
	FindOrderOutboxByID(ctx context.Context, id uint) (entity.OrderOutbox, error)
	DeleteOrderOutbox(context.Context, uint) error
}

type Service struct {
	publisher  broker.Publisher
	relayLimit int
	repo       Repository
}

func New(publisher broker.Publisher, repo Repository, relayLimit int) *Service {
	return &Service{
		publisher:  publisher,
		relayLimit: relayLimit,
		repo:       repo,
	}
}
