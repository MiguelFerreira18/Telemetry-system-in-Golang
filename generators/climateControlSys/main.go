package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
)

var global Generator
var isHealthy bool = true
var channel chan Telemetry

func init() {
	channel = make(chan Telemetry)
}

func main() {
	level := slog.LevelInfo
	if os.Getenv("APP_MODE") == "development" {
		level = slog.LevelDebug
	}

	global = initGenerator(level)
	global.Logger.Debug("CLIMATE CONTROL SYS GENERATOR IS ON")
	global.Logger.Debug("PID", "pid", os.Getpid())

	broker := initBroker()
	ctx := context.Background()
	err := broker.Connect(ctx)
	if err != nil {
		global.Logger.Error("An Error occurred while connecting to the broker", "broker", err)

	}
	global.Logger.Debug("Broker Initialized")
	go sendDataToMessageBroker(broker)

	global.Logger.Info("Climate control system Server Started")
	if err := router().ListenAndServe(); err != nil && err != http.ErrServerClosed {
		global.Logger.Error("Climate control system server failed", "error", err)
	}

}
