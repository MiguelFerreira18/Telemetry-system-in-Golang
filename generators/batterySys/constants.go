package main

const (
	ExitOk = 0
	ExitError = 1
)

const (
	HealthyVoltageBase  = 350.0
	HealthyVoltageRange = 70.0

	HealthyCurrentBase  = -10.0
	HealthyCurrentRange = 60.0

	HealthyTempBase  = 25.0
	HealthyTempRange = 15.0

	HealthySoCBase  = 60.0
	HealthySoCRange = 30.0

	HealthySoHBase  = 95.0
	HealthySoHRange = 5.0

	HealthyAuxVoltageBase  = 12.6
	HealthyAuxVoltageRange = 1.8
	HealthyAuxCurrentBase  = 5.0
	HealthyAuxCurrentRange = 25.0

	UnhealthyVoltageBase  = 280.0
	UnhealthyVoltageRange = 40.0

	UnhealthyCurrentBase  = 150.0
	UnhealthyCurrentRange = 100.0

	UnhealthyTempBase  = 65.0
	UnhealthyTempRange = 25.0

	UnhealthySoCBase  = 2.0
	UnhealthySoCRange = 8.0

	UnhealthySoHBase  = 60.0
	UnhealthySoHRange = 15.0

	UnhealthyAuxVoltageBase  = 10.5
	UnhealthyAuxVoltageRange = 1.5

	VoltageFaultOdds    = 5
	CurrentFaultOdds    = 4
	TempFaultOdds       = 3
	SoCFaultOdds        = 6
	SoHFaultOdds        = 10
	AuxVoltageFaultOdds = 7
)

const (
	ExchangeName = "telemetry"
	ExchangeType = "direct"

	QueueName = "Battery"
	RoutingKey = "Battery"
)
