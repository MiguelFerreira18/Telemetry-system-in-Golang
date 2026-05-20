package brokers

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Url     string
	conn    *amqp.Connection
	channel *amqp.Channel
}

// Close implements [MessageBroker].
func (r *Rabbit) Close() error {
	panic("unimplemented")
}

// Connect implements [MessageBroker].
func (r *Rabbit) Connect(ctx context.Context) error {
	conn, err := amqp.Dial(r.Url)
	if err != nil {
		return err
	}
	r.conn = conn
	r.channel, err = conn.Channel()
	if err != nil {
		return err
	}

	return r.declareExchange()
}

func (r *Rabbit) declareExchange() error {
	return r.channel.ExchangeDeclare(
		ExchangeName,
		ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
}

func (r *Rabbit) Subscribe(queue string, routingKey string, exchange string, callback func(msg []byte)) chan error {
	chanErr := make(chan error, 1)
	go func() {
		if err := r.queueDeclareAndBind(queue, routingKey); err != nil {
			chanErr <- fmt.Errorf("Unable to declare or bind to queue: %w", err)
			return
		}

		msgs, err := r.channel.Consume(queue, "", true, false, false, false, nil)
		if err != nil {
			chanErr <- fmt.Errorf("Unable to consume data from channel in %s queue with %s routing key: %w", queue, routingKey, err)
			return
		}

		for msg := range msgs {
			callback(msg.Body)
		}
	}()

	return chanErr
}

func (r *Rabbit) queueDeclareAndBind(queue string, routingKey string) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return r.channel.QueueBind(queue, routingKey, ExchangeName, false, nil)
}

type Mqtt struct {
	Url string
}

// Close implements [MessageBroker].
func (m *Mqtt) Close() error {
	panic("unimplemented")
}

// Connect implements [MessageBroker].
func (m *Mqtt) Connect(ctx context.Context) error {
	panic("unimplemented")
}

// Subscribe implements [MessageBroker].
func (m *Mqtt) Subscribe(queue string, routingKey string, exchange string, callback func(msg []byte)) chan error {
	panic("unimplemented")
}
