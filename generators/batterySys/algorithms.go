package main

import "math/rand"

type Healthy struct {
}

type Unhealthy struct {
}

func (h Healthy) GenerateVoltage() float32 {
	return HealthyVoltageBase + rand.Float32()*HealthyVoltageRange
}

func (h Healthy) GenerateCurrent() float32 {
	return HealthyCurrentBase + rand.Float32()*HealthyCurrentRange
}

func (h Healthy) GenerateTemp() float32 {
	return HealthyTempBase + rand.Float32()*HealthyTempRange
}

func (h Healthy) GenerateSoC() float32 {
	return HealthySoCBase + rand.Float32()*HealthySoCRange
}

func (h Healthy) GenerateSoH() float32 {
	return HealthySoHBase + rand.Float32()*HealthySoHRange
}

func (u Unhealthy) GenerateVoltage() float32 {
	if rand.Intn(VoltageFaultOdds) == 0 {
		return UnhealthyVoltageBase + rand.Float32()*UnhealthyVoltageRange
	}
	return HealthyVoltageBase + rand.Float32()*HealthyVoltageRange
}

func (u Unhealthy) GenerateCurrent() float32 {
	if rand.Intn(CurrentFaultOdds) == 0 {
		return UnhealthyCurrentBase + rand.Float32()*UnhealthyCurrentRange
	}
	return HealthyCurrentBase + rand.Float32()*HealthyCurrentRange
}

func (u Unhealthy) GenerateTemp() float32 {
	if rand.Intn(TempFaultOdds) == 0 {
		return UnhealthyTempBase + rand.Float32()*UnhealthyTempRange
	}
	return HealthyTempBase + rand.Float32()*HealthyTempRange
}

func (u Unhealthy) GenerateSoC() float32 {
	if rand.Intn(SoCFaultOdds) == 0 {
		return UnhealthySoCBase + rand.Float32()*UnhealthySoCRange
	}
	return HealthySoCBase + rand.Float32()*HealthySoCRange
}

func (u Unhealthy) GenerateSoH() float32 {
	if rand.Intn(SoHFaultOdds) == 0 {
		return UnhealthySoHBase + rand.Float32()*UnhealthySoHRange
	}
	return HealthySoHBase + rand.Float32()*HealthySoHRange
}
