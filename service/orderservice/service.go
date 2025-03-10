package orderservice

import (
	"context"
	"order/contract/broker"
	"order/entity"
)

type Repository interface {
	GetFirstOrderFromOutbox(ctx context.Context) (entity.OrderOutbox, error)
	UpdateOrderOutbox(ctx context.Context, orderOutbox entity.OrderOutbox) error
	CreateOrder(ctx context.Context, orderEntity entity.Order) (entity.Order, entity.OrderOutbox, error)
	CreateOrderOutbox(context.Context, entity.OrderOutbox) (entity.OrderOutbox, error)
	DeleteOrderOutbox(context.Context, uint) error
}

type Service struct {
	publisher broker.Publisher
	repo      Repository
}

func New(publisher broker.Publisher, repo Repository) *Service {
	return &Service{
		publisher: publisher,
		repo:      repo,
	}
}
