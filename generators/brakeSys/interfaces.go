package main

import "context"

type GenerationAlgo interface {
	GenerateBrakePressure() float32
	GenerateBrakeTemp() float32
	GeneratePadWear() float32
	GenerateAbsActive() bool
}

type MessageBroker interface {
	Connect(ctx context.Context) error
	Publish(ctx context.Context, topic string, payload []byte) error
	Close() error
}
