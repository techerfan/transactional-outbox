package shipmentservice

import (
	"context"
	"encoding/json"
	"fmt"
	"order/entity"
)

func (s *Service) ConsumeOrderEvent(msg []byte) error {
	var orderEvent entity.OrderEvent
	err := json.Unmarshal(msg, &orderEvent)
	if err != nil {
		return fmt.Errorf("could not unmarshal the msg: %v", err)
	}

	shipment := entity.Shipment{
		OrderID: orderEvent.Order.ID,
		Status:  entity.ShipmentStatusPending,
	}

	_, err = s.repo.CreateShipment(context.Background(), shipment)
	if err != nil {
		return fmt.Errorf("could not create the shipment: %v", err)
	}

	return nil
}
