package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"
)

func initGenerator(level slog.Level) Generator {
	opts := &slog.HandlerOptions{Level: level}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	generator := Generator{
		Logger: logger,
	}
	generator.setAlgo(Healthy{}, fuelType)
	return generator

}

type Generator struct {
	Logger    *slog.Logger
	algorithm GenerationAlgo
	mu        sync.Mutex
	cancel    context.CancelFunc
}

func (g *Generator) setAlgo(algo GenerationAlgo, fuelType int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.cancel != nil {
		g.cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())

	g.cancel = cancel
	g.algorithm = algo
	go g.runloop(ctx, fuelType)

}

func (g *Generator) runloop(ctx context.Context, fuel int) {
	g.Logger.Debug("Running Loop")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			g.Logger.Debug("Power train generator loop stopped")
			return
		case <-ticker.C:
			telemetry := g.generateTelemetry(fuel)
			for _, v := range telemetry {
				channel <- v
			}
			g.Logger.Debug("Generated telemetry", "count", len(telemetry))
		}
	}
}

func (g Generator) generateTelemetry(fuel int) []Telemetry {
	g.Logger.Debug("Algorithm type", "type", fmt.Sprintf("%T", g.algorithm))
	engineTemp := g.algorithm.GenerateEngineTemp()
	rpm := g.algorithm.GenerateRpm()
	fuelData := g.algorithm.GenerateFuelData(fuel)
	throttle := g.algorithm.GenerateThrottle()

	telemetry := make([]Telemetry, 5)
	for i := 0; i < len(telemetry); i++ {
		telemetry[i] = Telemetry{
			EngineTemp: engineTemp,
			Rpm:        rpm,
			Fuel:       fuelData,
			Throttle:   throttle,
		}
	}
	return telemetry
}
