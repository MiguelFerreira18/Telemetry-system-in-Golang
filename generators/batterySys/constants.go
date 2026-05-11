package main

const (
	ExitOk = 0
	ExitError = 1
)

const (
	// Healthy ranges
	HealthyVoltageBase  = 350.0
	HealthyVoltageRange = 70.0 // 350-420V

	HealthyCurrentBase  = -10.0
	HealthyCurrentRange = 60.0 // -10 to +50A

	HealthyTempBase  = 25.0
	HealthyTempRange = 15.0 // 25-40°C

	HealthySoCBase  = 60.0
	HealthySoCRange = 30.0 // 60-90%

	HealthySoHBase  = 95.0
	HealthySoHRange = 5.0 // 95-100%

	// Unhealthy ranges
	UnhealthyVoltageBase  = 280.0
	UnhealthyVoltageRange = 40.0 // 280-320V

	UnhealthyCurrentBase  = 150.0
	UnhealthyCurrentRange = 100.0 // 150-250A (Overcurrent)

	UnhealthyTempBase  = 65.0
	UnhealthyTempRange = 25.0 // 65-90°C (Overheating)

	UnhealthySoCBase  = 2.0
	UnhealthySoCRange = 8.0 // 2-10% (Critically Low)

	UnhealthySoHBase  = 60.0
	UnhealthySoHRange = 15.0 // 60-75% (Degraded)

	// Fault probabilities (1-in-N chance of unhealthy reading)
	VoltageFaultOdds = 5
	CurrentFaultOdds = 4
	TempFaultOdds    = 3
	SoCFaultOdds     = 6
	SoHFaultOdds     = 10
)

const (
	ExchangeName = "telemetry"
	ExchangeType = "direct"

	QueueName = "Battery"
	RoutingKey = "Battery"
)
