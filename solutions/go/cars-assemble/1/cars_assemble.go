package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(numberOfCars int, rate float64) float64 {
	return float64(numberOfCars) * rate / 100
}

func CalculateWorkingCarsPerMinute(numberOfCars int, rate float64) int {
	var carsHour float64 = CalculateWorkingCarsPerHour(numberOfCars, rate)
	return int(carsHour) / 60
}

func CalculateCost(numberOfCars int) uint {
	const cost = 10000
	const costOptimized = 95000
	return uint((numberOfCars/10)*costOptimized + ((numberOfCars % 10) * cost))
}
