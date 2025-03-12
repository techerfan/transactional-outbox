package orderservice

import (
	"context"
	"encoding/json"
	"fmt"
	"order/dto"
	"order/entity"
)

func (s *Service) PushToQueue(ctx context.Context, req dto.PushToQueueRequest) (dto.PushToQueueResponse, error) {
	orderOutboxes, err := s.repo.GetOrdersFromOutbox(ctx, s.relayLimit)
	if err != nil {
		return dto.PushToQueueResponse{}, fmt.Errorf("unexpected error while getting first order from outbox: %v", err)
	}

	for _, orderOutbox := range orderOutboxes {
		json, err := json.Marshal(orderOutbox)
		if err != nil {
			return dto.PushToQueueResponse{}, fmt.Errorf("unexpected error while marshalling order outbox item: %v", err)
		}

		err = s.publisher.Publish(entity.OrderCreated, string(json))
		if err != nil {
			return dto.PushToQueueResponse{}, fmt.Errorf("unexpected error while publishing order_created event: %v", err)
		}

		orderOutbox.Pushed = true
		err = s.repo.UpdateOrderOutbox(ctx, orderOutbox)
		if err != nil {
			return dto.PushToQueueResponse{}, fmt.Errorf("unexpected error while updating order outbox item: %v", err)
		}
	}

	return dto.PushToQueueResponse{}, nil
}
