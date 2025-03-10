package entity

type OrderOutbox struct {
	ID             uint   `json:"id"`
	OrderID        uint   `json:"order_id"`
	IdempotencyKey string `json:"idempotency_key"`
	Pushed         bool   `json:"pushed"`
}
