package rabbitmq

import (
	"fmt"

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
}

func New(config Config) (Adapter, error) {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		return Adapter{}, err
	}

	return Adapter{
		conn: conn,
	}, nil
}

func (a Adapter) Conn() *amqp.Connection {
	return a.conn
}
