package orderservice

import (
	"context"
	"order/dto"
	"order/entity"
)

func (s *Service) Submit(ctx context.Context, req dto.SubmitOrderRequest) (dto.SubmitOrderResponse, error) {
	// Create the order in the database
	_, _, err := s.repo.CreateOrder(ctx, entity.Order{
		Product: req.Prodcut,
		Amount:  req.Amount,
		Price:   req.Price,
	})
	if err != nil {
		return dto.SubmitOrderResponse{}, err
	}

	return dto.SubmitOrderResponse{}, nil
}
