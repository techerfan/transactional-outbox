package postgres

import (
	"order/entity"

	"gorm.io/gorm"
)

type (
	Order struct {
		gorm.Model

		Product string
		Amount  uint16
		Price   uint64
	}

	OrderOutbox struct {
		gorm.Model

		OrderID        uint
		IdempotencyKey string
		Pushed         bool
	}

	Shipment struct {
		gorm.Model

		OrderID uint
		Status  string
	}
)

func mapOrderToOrderEntity(order Order) entity.Order {
	return entity.Order{
		ID:      order.ID,
		Product: order.Product,
		Amount:  order.Amount,
		Price:   order.Price,
	}
}

func mapOrderEntityToOrder(order entity.Order) Order {
	return Order{
		Model: gorm.Model{
			ID: order.ID,
		},
		Product: order.Product,
		Amount:  order.Amount,
		Price:   order.Price,
	}
}

func mapOrderOutboxToOrderOutboxEntity(orderOutbox OrderOutbox) entity.OrderOutbox {
	return entity.OrderOutbox{
		ID:             orderOutbox.ID,
		OrderID:        orderOutbox.ID,
		IdempotencyKey: orderOutbox.IdempotencyKey,
		Pushed:         orderOutbox.Pushed,
	}
}

func mapOrderOutboxEntityToOrderOutbox(orderOutbox entity.OrderOutbox) OrderOutbox {
	return OrderOutbox{
		Model: gorm.Model{
			ID: orderOutbox.ID,
		},
		OrderID:        orderOutbox.OrderID,
		IdempotencyKey: orderOutbox.IdempotencyKey,
		Pushed:         orderOutbox.Pushed,
	}
}

func mapShipmentToShipmentEntity(shipment Shipment) entity.Shipment {
	return entity.Shipment{
		ID:        shipment.ID,
		OrderID:   shipment.OrderID,
		Status:    entity.ShipmentStatus(shipment.Status),
		CreatedAt: shipment.CreatedAt,
		UpdatedAt: shipment.UpdatedAt,
	}
}

func mapShipmentEntityToShipment(shipment entity.Shipment) Shipment {
	return Shipment{
		Model: gorm.Model{
			ID: shipment.ID,
		},
		OrderID: shipment.OrderID,
		Status:  string(shipment.Status),
	}
}
