package main

const (
	ExitOk = 0
	ExitError = 1
)

const (
	// Healthy ranges
	HealthyInteriorTempBase  = 18.0
	HealthyInteriorTempRange = 8.0 // 18-26°C

	HealthyExteriorTempBase  = 15.0
	HealthyExteriorTempRange = 20.0 // 15-35°C

	HealthyTargetTempBase  = 20.0
	HealthyTargetTempRange = 4.0 // 20-24°C

	HealthyFanSpeedBase  = 1
	HealthyFanSpeedRange = 5 // 1-6

	HealthyAirQualityBase  = 90.0
	HealthyAirQualityRange = 10.0 // 90-100%

	// Unhealthy ranges
	UnhealthyInteriorTempBase  = 35.0
	UnhealthyInteriorTempRange = 15.0 // 35-50°C (AC Failure)

	UnhealthyFanSpeedBase  = 0
	UnhealthyFanSpeedRange = 1 // 0 (Stuck)

	UnhealthyAirQualityBase  = 30.0
	UnhealthyAirQualityRange = 30.0 // 30-60% (Filter Issue)

	// Fault probabilities
	TempFaultOdds       = 5
	FanFaultOdds        = 8
	AirQualityFaultOdds = 6
)

const (
	ExchangeName = "telemetry"
	ExchangeType = "direct"

	QueueName = "ClimateControl"
	RoutingKey = "ClimateControl"
)
