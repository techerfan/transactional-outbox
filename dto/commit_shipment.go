package dto

type (
	CommitShipmentRequest struct {
		OrderID        uint   `json:"order_id"`
		OrderOutboxID  uint   `json:"order_outbox_id"`
		ShipmentID     uint   `json:"shipment_id"`
		IdempotencyKey string `json:"idempotency_key"`
	}

	CommitShipmentResponse struct{}
)
