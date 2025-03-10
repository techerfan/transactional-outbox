package shipment

import (
	"context"
	"fmt"
	"order/contract/goproto/shipment/shipment"
	"order/dto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	address string
}

func NewClient(address string) Client {
	return Client{address: address}
}

func (c Client) CommitShipment(ctx context.Context, req dto.CommitShipmentRequest) error {
	conn, err := grpc.NewClient(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("could not connect to the shipment service: %v", err)
	}
	defer conn.Close()

	client := shipment.NewShipmentServiceClient(conn)

	_, err = client.CommitShipment(ctx, &shipment.CommitShipmentRequest{
		Shipment: &shipment.Shipment{
			Id:             uint32(req.ShipmentID),
			OrderId:        uint32(req.OrderID),
			IdempotencyKey: req.IdempotencyKey,
		},
	})
	if err != nil {
		return fmt.Errorf("could not commit the shipment: %v", err)
	}

	return nil
}
