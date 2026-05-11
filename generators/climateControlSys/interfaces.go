package main

import "context"

type GenerationAlgo interface {
	GenerateInteriorTemp() float32
	GenerateExteriorTemp() float32
	GenerateTargetTemp() float32
	GenerateFanSpeed() int
	GenerateAirQuality() float32
}

type MessageBroker interface {
	Connect(ctx context.Context) error
	Publish(ctx context.Context, topic string, payload []byte) error
	Close() error
}
