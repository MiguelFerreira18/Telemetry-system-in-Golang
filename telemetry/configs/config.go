package configs

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"telemetry/internal/brokers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitSlog(appMode string) *slog.Logger {
	level := slog.LevelInfo
	if appMode == "development" {
		level = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{Level: level}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	return logger
}

func InitBroker(password string, user string, broker string) brokers.MessageBroker {
	url := ""
	if broker == "mqtt" {
		return &brokers.Mqtt{
			Url: url,
		}
	} else {
		host := os.Getenv("BROKER_HOST")
		if host == "" {
			host = "localhost"
		}
		url := fmt.Sprintf("amqp://%s:%s@%s:5672/", user, password, host)
		return &brokers.Rabbit{
			Url: url,
		}
	}
}

func InitServer() *http.Server {
	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
}
