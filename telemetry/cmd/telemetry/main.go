package main

import (
	"context"
	"net/http"
	"os"
	"telemetry/configs"
	"telemetry/internal/brokers"
	"telemetry/internal/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(metrics.EngineTemp)
	prometheus.MustRegister(metrics.Rpms)
	prometheus.MustRegister(metrics.RotationsTotal)
	prometheus.MustRegister(metrics.FuelAmount)
	prometheus.MustRegister(metrics.AngleAmount)
	prometheus.MustRegister(metrics.RateOfChange)
	prometheus.MustRegister(metrics.LogsInputTotal)
}

func main() {

	appMode := os.Getenv("APP_MODE")

	slog := configs.InitSlog(appMode)
	brokerPwd := os.Getenv("BROKER_PWD")
	brokerUser := os.Getenv("BROKER_USR")
	broker := os.Getenv("BROKER")
	messageBroker := configs.InitBroker(brokerPwd, brokerUser, broker)
	if err := messageBroker.Connect(context.Background()); err != nil {
		slog.Error("An error occurred while connecting to the message broker", "message_broker", "error", broker, err)
	}
	powerTrainProcessor := brokers.NewPowerTrainProcessor(slog)

	errorChannels := make([]chan error, 0)

	//Subscribe message brokers
	errorChannels = append(errorChannels, messageBroker.Subscribe(brokers.QueueName, brokers.RoutingKey, brokers.ExchangeName, powerTrainProcessor.PowerTrainHandler))

	for _, ch := range errorChannels {
		go func(errCh chan error) {
			if err := <-errCh; err != nil {
				slog.Error("Subscriber error", "error", err)
			}
		}(ch)
	}

	if err := configs.InitServer().ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("An error occurred while starting the server", "error", err)
	}

	slog.Info("Server closed successfully")
}
