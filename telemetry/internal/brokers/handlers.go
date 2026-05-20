package brokers

import (
	"encoding/json"
	"log/slog"
	"telemetry/internal/app"
	"telemetry/internal/metrics"
)

type PowertrainProcessor struct {
	slog *slog.Logger
}

func NewPowerTrainProcessor(slog *slog.Logger) PowertrainProcessor {
	return PowertrainProcessor{
		slog: slog,
	}
}

func (pp PowertrainProcessor) PowerTrainHandler(msg []byte) {
	var powerTrain app.PowertrainTelemetry
	if err := json.Unmarshal(msg, &powerTrain); err != nil {
		pp.slog.Error("An Error occurred while decoding message", "error", err)
	}
	// Process prometheus metrics
	metrics.EngineTemp.Add(float64(powerTrain.EngineTemp))
	metrics.Rpms.Add(float64(powerTrain.Rpm))
	metrics.RotationsTotal.Add(float64(powerTrain.Rpm))
	metrics.FuelAmount.Add(float64(powerTrain.Fuel.Quantity))
	metrics.AngleAmount.Add(float64(powerTrain.Throttle.Angle))
	metrics.RateOfChange.Add(float64(powerTrain.Throttle.RateOfChange))

	pp.slog.Info("Received Powertrain metric", "powertrain", powerTrain)
}
