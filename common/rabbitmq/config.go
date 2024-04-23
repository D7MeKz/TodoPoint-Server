package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

var (
	client  *amqp091.Connection
	channel *amqp091.Channel
)

func SetClient(url string) error {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return err
	}

	client = conn
	return nil
}

func GetClient() *amqp091.Connection {
	if client == nil {
		logrus.Fatal("Client does not exist")
	}
	return client
}

func SetChannel() error {
	client := GetClient()
	ch, err := client.Channel()
	if err != nil {
		return err
	}
	channel = ch
	return nil
}

func GetChannel() *amqp091.Channel {
	if channel == nil {
		logrus.Fatal("Channel does not exist")
	}
	return channel
}

func ConfigChannel(name string) error {
	channel = GetChannel()
	_, err := channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
