package main

const (
	FUEL = iota
	KW
)

const (
	ExitOk = 0
	ExitError = 1
)

const(
// Healthy ranges
	HealthyEngineTempBase  = 85
	HealthyEngineTempRange = 20 // 85–105°C

	HealthyRpmBase  = 1500
	HealthyRpmRange = 2500 // 1500–4000 RPM

	HealthyFuelBase  = 40.0
	HealthyFuelRange = 50.0 // 40–90%

	HealthyThrottleAngleBase  = 10
	HealthyThrottleAngleRange = 40 // 10–50°
	HealthyThrottleRocBase    = -5
	HealthyThrottleRocRange   = 10 // -5 to +5

	// Unhealthy ranges
	UnhealthyEngineTempBase  = 115
	UnhealthyEngineTempRange = 30 // 115–145°C

	UnhealthyRpmBase  = 5500
	UnhealthyRpmRange = 2000 // 5500–7500 RPM

	UnhealthyFuelRange = 8.0 // 0–8%

	UnhealthyThrottleAngleBase  = 70
	UnhealthyThrottleAngleRange = 20 // 70–90°
	UnhealthyThrottleRocBase    = 60
	UnhealthyThrottleRocRange   = 80 // 60–140

	// Fault probabilities (1-in-N chance of unhealthy reading)
	EngineTempFaultOdds = 3
	RpmFaultOdds        = 4
	FuelFaultOdds       = 3
	ThrottleFaultOdds   = 4
)

const (
	ExchangeName = "telemetry"
	ExchangeType = "direct"

	QueueName = "Powertrain"
	RoutingKey = "Powertrain"

)
