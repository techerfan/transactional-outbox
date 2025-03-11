package main

import (
	"order/adapter/rabbitmq"
	"order/adapter/shipment"
	"order/entity"
	"order/repository/postgres"
	"order/service/shipmentservice"
)

func main() {

	db := postgres.New(postgres.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
	})

	broker, err := rabbitmq.New(rabbitmq.Config{
		Host:     "localhost",
		Port:     5672,
		Username: "guest",
		Password: "guest",
	})
	if err != nil {
		panic(err)
	}

	client := shipment.NewClient(":8082")

	svc := shipmentservice.New(db, client)

	broker.Consume(entity.OrderCreated, svc.ConsumeOrderEvent)
}
