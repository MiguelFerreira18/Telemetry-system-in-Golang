package main

const (
	ExitOk = 0
	ExitError = 1
)
const (
	HealthyBrakePressureBase  = 0.0
	HealthyBrakePressureRange = 100.0

	HealthyBrakeTempBase  = 50.0
	HealthyBrakeTempRange = 150.0

	HealthyPadWearBase  = 0.0
	HealthyPadWearRange = 60.0

	UnhealthyBrakeTempBase  = 400.0
	UnhealthyBrakeTempRange = 200.0

	UnhealthyPadWearBase  = 85.0
	UnhealthyPadWearRange = 15.0

	UnhealthyPressureBase  = 0.0
	UnhealthyPressureRange = 20.0

	PressureFaultOdds    = 8
	TempFaultOdds        = 5
	WearFaultOdds        = 10
	AbsFaultOdds         = 6
	BrakeActiveFaultOdds = 4
)


const (
	ExchangeName = "telemetry"
	ExchangeType = "direct"

	QueueName = "Brake"
	RoutingKey = "Brake"
)
