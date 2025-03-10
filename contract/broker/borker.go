package broker

import "order/entity"

type MessageHandler func(msg []byte) error

type Publisher interface {
	Publish(event entity.Event, payload string) error
}

type Consumer interface {
	Consume(event entity.Event, handler MessageHandler) error
}
