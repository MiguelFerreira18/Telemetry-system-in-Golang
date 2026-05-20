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
			g.Logger.Debug("Battery generator loop stopped")
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
	voltage := g.algorithm.GenerateVoltage()
	current := g.algorithm.GenerateCurrent()
	temp := g.algorithm.GenerateTemp()
	soc := g.algorithm.GenerateSoC()
	soh := g.algorithm.GenerateSoH()
	auxV := g.algorithm.GenerateAuxVoltage()
	auxC := g.algorithm.GenerateAuxCurrent()
	starter := g.algorithm.GenerateStarterStatus()
	alt := g.algorithm.GenerateAlternatorLoad()

	telemetry := make([]Telemetry, 5)
	for i := range telemetry {
		telemetry[i] = Telemetry{
			Voltage:        voltage,
			Current:        current,
			Temperature:    temp,
			SoC:            soc,
			SoH:            soh,
			AuxVoltage:     auxV,
			AuxCurrent:     auxC,
			StarterStatus:  starter,
			AlternatorLoad: alt,
		}
	}
	return telemetry
}
