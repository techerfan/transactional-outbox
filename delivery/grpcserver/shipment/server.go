package shipment

import (
	"context"
	"order/contract/goproto/shipment/shipment"
	"order/service/orderservice"
)

type Server struct {
	shipment.UnimplementedShipmentServiceServer
	svc orderservice.Service
}

func New(svc orderservice.Service) *Server {
	return &Server{
		UnimplementedShipmentServiceServer: shipment.UnimplementedShipmentServiceServer{},
		svc:                                svc,
	}
}

func (s *Server) CommitShipment(ctx context.Context, req *shipment.CommitShipmentRequest) (*shipment.CommitShipmentResponse, error) {
	return &shipment.CommitShipmentResponse{
		Message: "Shipment committed",
	}, nil
}
