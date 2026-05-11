package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var global Generator
var fuelType int
var isHealthy bool = true
var channel chan Telemetry

func init() {
	channel = make(chan Telemetry)
	fuelType = FUEL
}

func main() {
	level := slog.LevelInfo
	if os.Getenv("APP_MODE") == "development" {
		level = slog.LevelDebug
	}

	global = initGenerator(level)
	global.Logger.Debug("POWER TRAIN SYS GENERATOR IS ON")
	global.Logger.Debug("PID", "pid", os.Getpid())

	broker := initBroker()
	ctx := context.Background()
	err := broker.Connect(ctx)
	if err != nil {
		global.Logger.Error("An Error occurred while connecting to the broker", "broker", err)

	}
	global.Logger.Debug("Broker Initialized")
	go sendDataToMessageBroker(broker)

	global.Logger.Info("Power train system Server Started")
	if err := router().ListenAndServe(); err != nil && err != http.ErrServerClosed {
		global.Logger.Error("Power train system server failed", "error", err)
	}

}

func initBroker() MessageBroker {
	password := os.Getenv("BROKER_PWD")
	user := os.Getenv("BROKER_USR")
	url := ""

	if os.Getenv("BROKER") == "mqtt" {
		return &Mqtt{
			url: url,
		}
	} else {
		url = fmt.Sprintf("amqp://%s:%s@localhost:5672/", user, password)
		return &Rabbit{
			url: url,
		}
	}
}

func sendDataToMessageBroker(broker MessageBroker) {
	ctx := context.Background()
	for v := range channel {
		body, err := json.Marshal(v)
		if err != nil {
			global.Logger.Error("Failed to marshal telemetry", "Marshal Telemetry", err)
			continue
		}
		if err := broker.Publish(ctx, "Powertrain", body); err != nil {
			global.Logger.Error("Failed to publish telemetry", "Publish Telemetry", err)
		}
	}
}

func router() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if !isHealthy {
				global.setAlgo(Healthy{}, fuelType)
				isHealthy = true
				global.Logger.Info("Power train System generator switched to healthy mode")
			} else {
				global.Logger.Info("System already in healthy mode")
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	router.HandleFunc("/unhealthy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if isHealthy {
				global.setAlgo(Unhealthy{}, fuelType)
				isHealthy = false
				global.Logger.Info("Power train System generator switched to unhealthy mode")
			}
			w.WriteHeader(http.StatusOK)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	router.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			global.Logger.Info("System Shutting Down")
			os.Exit(ExitOk)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	return &http.Server{
		Addr:    ":9091",
		Handler: router,
	}
}

func initGenerator(level slog.Level) Generator {
	opts := &slog.HandlerOptions{Level: level}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	generator := Generator{
		Logger: logger,
	}
	generator.setAlgo(Healthy{}, fuelType)
	return generator

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
