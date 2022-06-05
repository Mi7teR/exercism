package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate) / 100 * successRate / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	cost := 10000
	costOfTenCarsGroup := 95000
	tenCarsGroup := carsCount / 10
	summaryCost := tenCarsGroup*costOfTenCarsGroup + (carsCount-tenCarsGroup*10)*cost

	return uint(summaryCost)
}
