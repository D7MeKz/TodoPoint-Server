package rabbitmq

import "github.com/rabbitmq/amqp091-go"

func GetClient(url string) (*amqp091.Connection, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
