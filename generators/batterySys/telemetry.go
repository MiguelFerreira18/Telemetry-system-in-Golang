package main

type Telemetry struct {
	Voltage     float32
	Current     float32
	Temperature float32
	SoC         float32
	SoH         float32

	AuxVoltage     float32
	AuxCurrent     float32
	StarterStatus  bool
	AlternatorLoad float32
}
