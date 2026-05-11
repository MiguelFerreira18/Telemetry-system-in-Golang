package main

type Telemetry struct {
	EngineTemp int
	Rpm        int
	Fuel       FuelData
	Throttle   ThrottlePosition
}

type FuelData struct {
	FuelType int
	Quantity float32
}

type ThrottlePosition struct {
	Angle        uint8
	RateOfChange int
}
