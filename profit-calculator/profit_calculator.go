package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64 = 10

	fmt.Print("Enter Revenue: ")
	fmt.Scanf("%f\n", &revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scanf("%f\n", &expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scanf("%f\n", &taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earning Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
