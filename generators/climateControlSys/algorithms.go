package main

import "math/rand"

type Healthy struct {
}

type Unhealthy struct {
}

func (h Healthy) GenerateInteriorTemp() float32 {
	return HealthyInteriorTempBase + rand.Float32()*HealthyInteriorTempRange
}

func (h Healthy) GenerateExteriorTemp() float32 {
	return HealthyExteriorTempBase + rand.Float32()*HealthyExteriorTempRange
}

func (h Healthy) GenerateTargetTemp() float32 {
	return HealthyTargetTempBase + rand.Float32()*HealthyTargetTempRange
}

func (h Healthy) GenerateFanSpeed() int {
	return HealthyFanSpeedBase + rand.Intn(HealthyFanSpeedRange)
}

func (h Healthy) GenerateAirQuality() float32 {
	return HealthyAirQualityBase + rand.Float32()*HealthyAirQualityRange
}

func (u Unhealthy) GenerateInteriorTemp() float32 {
	if rand.Intn(TempFaultOdds) == 0 {
		return UnhealthyInteriorTempBase + rand.Float32()*UnhealthyInteriorTempRange
	}
	return HealthyInteriorTempBase + rand.Float32()*HealthyInteriorTempRange
}

func (u Unhealthy) GenerateExteriorTemp() float32 {
	return HealthyExteriorTempBase + rand.Float32()*HealthyExteriorTempRange
}

func (u Unhealthy) GenerateTargetTemp() float32 {
	return HealthyTargetTempBase + rand.Float32()*HealthyTargetTempRange
}

func (u Unhealthy) GenerateFanSpeed() int {
	if rand.Intn(FanFaultOdds) == 0 {
		return UnhealthyFanSpeedBase
	}
	return HealthyFanSpeedBase + rand.Intn(HealthyFanSpeedRange)
}

func (u Unhealthy) GenerateAirQuality() float32 {
	if rand.Intn(AirQualityFaultOdds) == 0 {
		return UnhealthyAirQualityBase + rand.Float32()*UnhealthyAirQualityRange
	}
	return HealthyAirQualityBase + rand.Float32()*HealthyAirQualityRange
}
