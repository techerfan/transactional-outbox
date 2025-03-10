package shipmentservice

import (
	"context"
	"order/entity"
)

type Repository interface {
	CreateShipment(ctx context.Context, shipment entity.Shipment) (entity.Shipment, error)
	UpdateShipment(ctx context.Context, shipment entity.Shipment) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
