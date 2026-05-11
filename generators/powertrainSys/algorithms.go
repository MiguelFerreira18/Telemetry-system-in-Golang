package main

import "math/rand"

type Healthy struct {
}

type Unhealthy struct {
}

func (h Healthy) GenerateEngineTemp() int {
	return HealthyEngineTempBase + rand.Intn(HealthyEngineTempRange)
}
func (h Healthy) GenerateRpm() int {
	return HealthyRpmBase + rand.Intn(HealthyRpmRange)

}
func (h Healthy) GenerateFuelData(fuelType int) FuelData {
	return FuelData{
		FuelType: fuelType,
		Quantity: HealthyFuelBase + rand.Float32()*HealthyFuelRange,
	}

}
func (h Healthy) GenerateThrottle() ThrottlePosition {
	return ThrottlePosition{
		Angle:        HealthyThrottleAngleBase + uint8(rand.Intn(HealthyThrottleAngleRange)),
		RateOfChange: HealthyThrottleRocBase + rand.Intn(HealthyThrottleRocRange),
	}

}
func (u Unhealthy) GenerateEngineTemp() int {
	if rand.Intn(EngineTempFaultOdds) == 0 {
		return UnhealthyEngineTempBase + rand.Intn(UnhealthyEngineTempRange)
	}
	return HealthyEngineTempBase + rand.Intn(HealthyEngineTempRange)

}
func (u Unhealthy) GenerateRpm() int {
	if rand.Intn(RpmFaultOdds) == 0 {
		return UnhealthyRpmBase + rand.Intn(UnhealthyRpmRange)
	}
	return HealthyRpmBase + rand.Intn(HealthyRpmRange)
}
func (u Unhealthy) GenerateFuelData(fuelType int) FuelData {
	qty := HealthyFuelBase + rand.Float32()*HealthyFuelRange
	if rand.Intn(FuelFaultOdds) == 0 {
		qty = rand.Float32() * UnhealthyFuelRange
	}
	return FuelData{
		FuelType: fuelType,
		Quantity: qty,
	}

}
func (u Unhealthy) GenerateThrottle() ThrottlePosition {
	if rand.Intn(ThrottleFaultOdds) == 0 {
		return ThrottlePosition{
			Angle:        UnhealthyThrottleAngleBase + uint8(rand.Intn(UnhealthyThrottleAngleRange)),
			RateOfChange: UnhealthyThrottleRocBase + rand.Intn(UnhealthyThrottleRocRange),
		}
	}
	return ThrottlePosition{
		Angle:        HealthyThrottleAngleBase + uint8(rand.Intn(HealthyThrottleAngleRange)),
		RateOfChange: HealthyThrottleRocBase + rand.Intn(HealthyThrottleRocRange),
	}
}
