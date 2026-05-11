package main

import "context"

type GenerationAlgo interface {
	GenerateVoltage() float32
	GenerateCurrent() float32
	GenerateTemp() float32
	GenerateSoC() float32
	GenerateSoH() float32
	GenerateAuxVoltage() float32
	GenerateAuxCurrent() float32
	GenerateStarterStatus() bool
	GenerateAlternatorLoad() float32
}

type MessageBroker interface {
	Connect(ctx context.Context) error
	Publish(ctx context.Context, topic string, payload []byte) error
	Close() error
}
