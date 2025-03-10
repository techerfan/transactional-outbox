package entity

import "time"

type ShipmentStatus string

const (
	ShipmentStatusPending   ShipmentStatus = "pending"
	ShipmentStatusShipped   ShipmentStatus = "shipped"
	ShipmentStatusDelivered ShipmentStatus = "delivered"
)

type Shipment struct {
	ID        uint           `json:"id"`
	OrderID   uint           `json:"order_id"`
	Status    ShipmentStatus `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
