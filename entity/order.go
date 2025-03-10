package entity

type Order struct {
	ID      uint   `json:"id"`
	Product string `json:"product"`
	Amount  uint16 `json:"amount"`
	Price   uint64 `json:"price"`
}
