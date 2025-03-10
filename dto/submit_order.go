package dto

type (
	SubmitOrderRequest struct {
		Prodcut string
		Amount  uint16
		Price   uint64
	}

	SubmitOrderResponse struct{}
)
