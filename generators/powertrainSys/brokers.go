package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func initBroker() MessageBroker {
	password := os.Getenv("BROKER_PWD")
	user := os.Getenv("BROKER_USR")
	url := ""

	if os.Getenv("BROKER") == "mqtt" {
		return &Mqtt{
			url: url,
		}
	} else {
		url = fmt.Sprintf("amqp://%s:%s@localhost:5672/", user, password)
		return &Rabbit{
			url: url,
		}
	}
}

type Rabbit struct {
	url     string
	conn    *amqp.Connection
	channel *amqp.Channel
}

type Mqtt struct {
	url string
}

func (r *Rabbit) Connect(ctx context.Context) error {
	conn, err := amqp.Dial(r.url)
	if err != nil {
		return err
	}
	r.conn = conn
	r.channel, err = conn.Channel()
	if err != nil {
		return err
	}

	if err := r.declareExchange(); err != nil {
		return err
	}
	if err := r.queueDeclare(); err != nil {
		return err
	}

	return nil
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
func (r *Rabbit) queueDeclare() error {
	_, err := r.channel.QueueDeclare(QueueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return r.channel.QueueBind(QueueName, RoutingKey, ExchangeName, false, nil)
}

func (r *Rabbit) Publish(ctx context.Context, topic string, payload []byte) error {
	return r.channel.PublishWithContext(ctx,
		"",
		topic,
		false, false, amqp.Publishing{Body: payload})
}

func (r *Rabbit) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.conn.Close()

}

func (r Mqtt) Connect(ctx context.Context) error {

	return nil
}

func (m Mqtt) Publish(ctx context.Context, topic string, payload []byte) error {
	var telemetry Telemetry
	err := json.Unmarshal(payload, &telemetry)
	if err != nil {
		return err
	}
	global.Logger.Debug("Telemetry Unmarshal", "telemetry data", telemetry)
	return nil
}

func (r Mqtt) Close() error {

	return nil
}

func sendDataToMessageBroker(broker MessageBroker) {
	ctx := context.Background()
	for v := range channel {
		body, err := json.Marshal(v)
		if err != nil {
			global.Logger.Error("Failed to marshal telemetry", "Marshal Telemetry", err)
			continue
		}
		if err := broker.Publish(ctx, "Powertrain", body); err != nil {
			global.Logger.Error("Failed to publish telemetry", "Publish Telemetry", err)
		}
	}
}
