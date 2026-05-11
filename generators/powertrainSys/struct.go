package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"math/rand"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Generator struct {
	Logger    *slog.Logger
	algorithm GenerationAlgo
	mu        sync.Mutex
	cancel    context.CancelFunc
}

type Rabbit struct {
	url     string
	conn    *amqp.Connection
	channel *amqp.Channel
}

type Mqtt struct {
	url string
}

type Telemetry struct {
	EngineTemp int
	Rpm        int
	Fuel       FuelData
	Throttle   ThrottlePosition
}

type FuelData struct {
	FuelType int
	Quantity float32
}

type ThrottlePosition struct {
	Angle        uint8
	RateOfChange int
}

type Healthy struct {
}

type Unhealthy struct {
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

func (h Healthy) GenerateEngineTemp() int {
	return HealthyEngineTempBase + rand.Intn(HealthyEngineTempRange)
}
func (h Healthy) GenerateRpm() int {
	return HealthyRpmBase + rand.Intn(HealthyRpmRange)

}
func (h Healthy) GenerateFuelData(fuelType int) FuelData {
	return FuelData{
		FuelType: fuelType,
		Quantity: HealthyFuelBase + rand.Float32()*HealthyFuelRange,
	}

}
func (h Healthy) GenerateThrottle() ThrottlePosition {
	return ThrottlePosition{
		Angle:        HealthyThrottleAngleBase + uint8(rand.Intn(HealthyThrottleAngleRange)),
		RateOfChange: HealthyThrottleRocBase + rand.Intn(HealthyThrottleRocRange),
	}

}
func (u Unhealthy) GenerateEngineTemp() int {
	if rand.Intn(EngineTempFaultOdds) == 0 {
		return UnhealthyEngineTempBase + rand.Intn(UnhealthyEngineTempRange)
	}
	return HealthyEngineTempBase + rand.Intn(HealthyEngineTempRange)

}
func (u Unhealthy) GenerateRpm() int {
	if rand.Intn(RpmFaultOdds) == 0 {
		return UnhealthyRpmBase + rand.Intn(UnhealthyRpmRange)
	}
	return HealthyRpmBase + rand.Intn(HealthyRpmRange)
}
func (u Unhealthy) GenerateFuelData(fuelType int) FuelData {
	qty := HealthyFuelBase + rand.Float32()*HealthyFuelRange
	if rand.Intn(FuelFaultOdds) == 0 {
		qty = rand.Float32() * UnhealthyFuelRange
	}
	return FuelData{
		FuelType: fuelType,
		Quantity: qty,
	}

}
func (u Unhealthy) GenerateThrottle() ThrottlePosition {
	if rand.Intn(ThrottleFaultOdds) == 0 {
		return ThrottlePosition{
			Angle:        UnhealthyThrottleAngleBase + uint8(rand.Intn(UnhealthyThrottleAngleRange)),
			RateOfChange: UnhealthyThrottleRocBase + rand.Intn(UnhealthyThrottleRocRange),
		}
	}
	return ThrottlePosition{
		Angle:        HealthyThrottleAngleBase + uint8(rand.Intn(HealthyThrottleAngleRange)),
		RateOfChange: HealthyThrottleRocBase + rand.Intn(HealthyThrottleRocRange),
	}
}

// BROKERS
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
