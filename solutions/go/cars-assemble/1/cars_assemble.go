package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    workingCarsPerHour := float64(productionRate) * successRate / 100
    
    return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    productionPerMinute := int(productionPerHour) / 60
    return productionPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupTen := carsCount / 10
    notGroup := carsCount % 10

    result := (groupTen * 95000) + (notGroup * 10000)
    return uint(result)
}
