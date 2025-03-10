package shipment

import (
	"context"
	"log"
	"net"
	"order/contract/goproto/shipment/shipment"
	"order/dto"
	"order/service/orderservice"

	"google.golang.org/grpc"
)

type Server struct {
	shipment.UnimplementedShipmentServiceServer
	svc orderservice.Service
}

func New(svc orderservice.Service) Server {
	return Server{
		UnimplementedShipmentServiceServer: shipment.UnimplementedShipmentServiceServer{},
		svc:                                svc,
	}
}

func (s Server) CommitShipment(ctx context.Context, req *shipment.CommitShipmentRequest) (*shipment.CommitShipmentResponse, error) {
	_, err := s.svc.CommitShipment(ctx, dto.CommitShipmentRequest{
		OrderID:        uint(req.Shipment.OrderId),
		OrderOutboxID:  uint(req.Shipment.OrderOutboxId),
		ShipmentID:     uint(req.Shipment.Id),
		IdempotencyKey: req.Shipment.IdempotencyKey,
	})
	if err != nil {
		return nil, err
	}

	return &shipment.CommitShipmentResponse{
		Message: "Shipment committed",
	}, nil
}

func (s Server) Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	grpcServer := grpc.NewServer()

	shipment.RegisterShipmentServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
