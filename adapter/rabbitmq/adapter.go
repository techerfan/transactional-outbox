package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"order/contract/broker"
	"order/entity"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func (a Adapter) Publish(event entity.Event, payload string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	ch, err := a.conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		string(event),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("could not declare the queue: %v", err)
	}

	err = ch.PublishWithContext(ctx,
		"",
		string(event),
		true, // mandatory
		false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(payload),
		},
	)
	if err != nil {
		return fmt.Errorf("could not publish the msg: %v", err)
	}

	return nil
}

func (a Adapter) Consume(event entity.Event, handler broker.MessageHandler) error {
	ch, err := a.conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		string(event),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("could not declare the queue: %v", err)
	}

	msgs, err := ch.Consume(
		string(event),
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("could not consume the msg: %v", err)
	}

	for msg := range msgs {
		time.Sleep(time.Second * 1)
		err := handler(msg.Body)
		if err != nil {
			msg.Nack(false, true)
			log.Printf("could not handle the msg: %v, msg: %s", err, msg.Body)
			continue
		}
		msg.Ack(false)
	}

	return nil
}
