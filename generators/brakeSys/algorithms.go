package main

import "math/rand"

type Healthy struct {
}

type Unhealthy struct {
}

func (h Healthy) GenerateBrakePressure() float32 {
	return HealthyBrakePressureBase + rand.Float32()*HealthyBrakePressureRange
}

func (h Healthy) GenerateBrakeTemp() float32 {
	return HealthyBrakeTempBase + rand.Float32()*HealthyBrakeTempRange
}

func (h Healthy) GeneratePadWear() float32 {
	return HealthyPadWearBase + rand.Float32()*HealthyPadWearRange
}

func (h Healthy) GenerateAbsActive() bool {
	return rand.Intn(10) == 0
}


func (u Unhealthy) GenerateBrakePressure() float32 {
	if rand.Intn(PressureFaultOdds) == 0 {
		return UnhealthyPressureBase + rand.Float32()*UnhealthyPressureRange
	}
	return HealthyBrakePressureBase + rand.Float32()*HealthyBrakePressureRange
}

func (u Unhealthy) GenerateBrakeTemp() float32 {
	if rand.Intn(TempFaultOdds) == 0 {
		return UnhealthyBrakeTempBase + rand.Float32()*UnhealthyBrakeTempRange
	}
	return HealthyBrakeTempBase + rand.Float32()*HealthyBrakeTempRange
}

func (u Unhealthy) GeneratePadWear() float32 {
	if rand.Intn(WearFaultOdds) == 0 {
		return UnhealthyPadWearBase + rand.Float32()*UnhealthyPadWearRange
	}
	return HealthyPadWearBase + rand.Float32()*HealthyPadWearRange
}

func (u Unhealthy) GenerateAbsActive() bool {
	if rand.Intn(AbsFaultOdds) == 0 {
		return true
	}
	return rand.Intn(5) == 0
}
