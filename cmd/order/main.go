package main

import (
	"order/adapter/rabbitmq"
	"order/delivery/grpcserver/shipment"
	"order/delivery/httpserver"
	"order/repository/postgres"
	"order/service/orderservice"
)

func main() {

	rabbitmqAdapter, err := rabbitmq.New(rabbitmq.Config{
		Username: "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     5672,
	})
	if err != nil {
		panic(err)
	}

	db := postgres.New(postgres.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
	})

	orderService := orderservice.New(rabbitmqAdapter, db)

	grpcServer := shipment.New(*orderService)
	go func() {
		grpcServer.Start(":8082")
	}()

	server := httpserver.New(orderService)
	server.Start(":8081")
}
