package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2

func main() {
	fmt.Println("Hello, World!")

	var investmentAmount, expectedReturnRate float64
	var years float64 = 5

	fmt.Print("Enter Invested Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Investment Period: ")
	fmt.Scan(&years)

	futureValue, realfutureValue := calculateInvestments(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future Value: %.1f\n", futureValue)
	fmt.Printf("Future Value(adjusted for Inflation): %.2f\n", realfutureValue)
}

func calculateInvestments(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
