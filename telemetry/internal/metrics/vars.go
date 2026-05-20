package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	EngineTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "engine_temperature_c",
		Help: "Engine temperature in motor in celcius",
	})
	Rpms = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "rotations_per_minutes",
		Help: "Rotation per minute of the engine shaft",
	})
	RotationsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "rotations_total",
		Help: "Total rotations during entire operation of engine shaft",
	})

	//Quantity
	FuelAmount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "fuel_amount_liters",
		Help: "Amount of fuel in the tanks in liters",
	})

	AngleAmount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "angle_amount_degrees",
		Help: "Amount of degrees in rotation or angle calculation.",
	})

	RateOfChange = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "rate_of_change_throttle_position",
		Help: "Rate of change of throttle position over time",
	})

	LogsInputTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "logs_input_total",
		Help: "Total amount of logs",
	})
)
