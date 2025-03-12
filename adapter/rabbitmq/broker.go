package rabbitmq

import (
	"fmt"
	"order/entity"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
}

type Adapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New(config Config) (Adapter, error) {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		return Adapter{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return Adapter{}, err
	}

	_, err = ch.QueueDeclare(
		string(entity.OrderCreated),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return Adapter{}, fmt.Errorf("could not declare the queue: %v", err)
	}

	return Adapter{
		conn: conn,
		ch:   ch,
	}, nil
}

func (a Adapter) Conn() *amqp.Connection {
	return a.conn
}
