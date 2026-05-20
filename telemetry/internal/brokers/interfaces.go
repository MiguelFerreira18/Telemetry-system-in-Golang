package brokers

import "context"

type MessageBroker interface {
	Connect(ctx context.Context) error
	Subscribe(queue string, routingKey string, exchange string, callback func(msg []byte)) chan error
	Close() error
}
