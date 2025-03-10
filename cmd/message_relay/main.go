package main

import (
	"context"
	"log"
	"order/adapter/rabbitmq"
	"order/repository/postgres"
	"order/scheduler"
	"order/service/orderservice"
	"sync"
)

func main() {
	borker, err := rabbitmq.New(rabbitmq.Config{
		Username: "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     5672,
	})
	if err != nil {
		log.Fatalf("could not create the broker: %v", err)
	}

	repo := postgres.New(postgres.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
	})

	orderservice := orderservice.New(borker, repo)

	scheduler := scheduler.New(scheduler.Config{
		SendOrdersToOutboxInterval: 10,
	}, *orderservice)

	ctx := context.Background()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	scheduler.Start(ctx, wg)

	wg.Wait()
}
