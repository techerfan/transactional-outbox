package entity

type OrderEvent struct {
	Order          Order  `json:"order"`
	IdempotencyKey string `json:"idempotency_key"`
}
