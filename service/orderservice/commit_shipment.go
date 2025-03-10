package orderservice

import (
	"context"
	"errors"
	"order/dto"
)

func (s *Service) CommitShipment(ctx context.Context, req dto.CommitShipmentRequest) (dto.CommitShipmentResponse, error) {
	// TODO: validations must be done in the validator layer
	_, err := s.repo.FindOrderByID(ctx, req.OrderID)
	if err != nil {
		return dto.CommitShipmentResponse{}, err
	}

	orderOutbox, err := s.repo.FindOrderOutboxByID(ctx, req.OrderOutboxID)
	if err != nil {
		return dto.CommitShipmentResponse{}, err
	}

	if !orderOutbox.Pushed {
		return dto.CommitShipmentResponse{}, errors.New("order is not pushed yet")
	}

	if orderOutbox.IdempotencyKey != req.IdempotencyKey {
		return dto.CommitShipmentResponse{}, errors.New("idempotency key does not match")
	}

	if err := s.repo.DeleteOrderOutbox(ctx, orderOutbox.ID); err != nil {
		return dto.CommitShipmentResponse{}, err
	}

	return dto.CommitShipmentResponse{}, nil
}
