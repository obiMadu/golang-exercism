package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate / 100.0) * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((CalculateWorkingCarsPerHour(productionRate, successRate)) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// CostofEachTenCars contains the total amount for each group of ten cars
	costOfEachTenCars := (carsCount / 10) * 95000

	// CostofTheRemainingCars contains the total amount for other cars after the last group of 10
	costofTheRemainingCars := (carsCount % 10) * 10000

	return uint(costOfEachTenCars + costofTheRemainingCars)
}
