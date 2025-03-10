package shipmentservice

import (
	"context"
	"encoding/json"
	"fmt"
	"order/dto"
	"order/entity"
)

func (s *Service) ConsumeOrderEvent(msg []byte) error {
	var orderEvent entity.OrderEvent
	err := json.Unmarshal(msg, &orderEvent)
	if err != nil {
		return fmt.Errorf("could not unmarshal the msg: %v", err)
	}

	// check if the order is already committed
	shipment, exist, err := s.repo.FindShipmentByOrderID(context.Background(), orderEvent.Order.ID)
	if err != nil {
		return fmt.Errorf("could not find the shipment: %v", err)
	}

	// Create a new shipment if it doesn't exist
	if !exist {
		shipment = entity.Shipment{
			OrderID: orderEvent.Order.ID,
			Status:  entity.ShipmentStatusPending,
		}

		_, err = s.repo.CreateShipment(context.Background(), shipment)
		if err != nil {
			return fmt.Errorf("could not create the shipment: %v", err)
		}
	}

	err = s.shipmentClient.CommitShipment(context.Background(), dto.CommitShipmentRequest{
		ShipmentID:     shipment.ID,
		OrderID:        shipment.OrderID,
		IdempotencyKey: orderEvent.IdempotencyKey,
	})
	if err != nil {
		return fmt.Errorf("could not commit the shipment: %v", err)
	}

	return nil
}
