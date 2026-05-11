package main

import "context"

type GenerationAlgo interface {
	GenerateEngineTemp() int
	GenerateRpm() int
	GenerateFuelData(fuel int) FuelData
	GenerateThrottle() ThrottlePosition
}

type MessageBroker interface {
	Connect(ctx context.Context) error
	Publish(ctx context.Context, topic string, payload []byte) error
	Close() error
}
