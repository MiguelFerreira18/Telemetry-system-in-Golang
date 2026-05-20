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
	generator.setAlgo(Healthy{}, 0)
	return generator

}

type Generator struct {
	Logger    *slog.Logger
	algorithm GenerationAlgo
	mu        sync.Mutex
	cancel    context.CancelFunc
}

func (g *Generator) setAlgo(algo GenerationAlgo, _ int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.cancel != nil {
		g.cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())

	g.cancel = cancel
	g.algorithm = algo
	go g.runloop(ctx)

}

func (g *Generator) runloop(ctx context.Context) {
	g.Logger.Debug("Running Loop")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			g.Logger.Debug("Brake generator loop stopped")
			return
		case <-ticker.C:
			telemetry := g.generateTelemetry()
			for _, v := range telemetry {
				channel <- v
			}
			g.Logger.Debug("Generated telemetry", "count", len(telemetry))
		}
	}
}

func (g Generator) generateTelemetry() []Telemetry {
	g.Logger.Debug("Algorithm type", "type", fmt.Sprintf("%T", g.algorithm))
	pressure := g.algorithm.GenerateBrakePressure()
	temp := g.algorithm.GenerateBrakeTemp()
	wear := g.algorithm.GeneratePadWear()
	abs := g.algorithm.GenerateAbsActive()

	telemetry := make([]Telemetry, 5)
	for i := range telemetry {
		telemetry[i] = Telemetry{
			BrakePressure: pressure,
			BrakeTemp:     temp,
			PadWear:       wear,
			AbsActive:     abs,
		}
	}
	return telemetry
}
