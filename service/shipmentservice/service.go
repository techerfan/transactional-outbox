package shipmentservice

import (
	"context"
	"order/dto"
	"order/entity"
)

type Repository interface {
	CreateShipment(ctx context.Context, shipment entity.Shipment) (entity.Shipment, error)
	UpdateShipment(ctx context.Context, shipment entity.Shipment) error
	FindShipmentByOrderID(ctx context.Context, orderID uint) (entity.Shipment, bool, error)
}

type shipmentClient interface {
	CommitShipment(ctx context.Context, req dto.CommitShipmentRequest) error
}

type Service struct {
	repo           Repository
	shipmentClient shipmentClient
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
