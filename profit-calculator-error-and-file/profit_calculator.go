package main

import (
	"errors"
	"fmt"
	"os"
)

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

	isValidRevenue, err := validateInput(revenue)

	if err != nil {
		fmt.Println("Error: Invalid Revenue")
	}

	isValidExpenses, err := validateInput(expenses)

	if err != nil {
		fmt.Println("Error: Invalid Expenses")
	}

	isValidTaxRate, err := validateInput(taxRate)

	if err != nil {
		fmt.Println("Error: Invalid Tax Rate")
	}

	if isValidRevenue && isValidExpenses && isValidTaxRate {
		ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

		outputToStore := fmt.Sprintf(
			`Earning Before Tax: %.1f
		Profit: %.1f
		Ratio: %.1f
		`, ebt, profit, ratio)

		fmt.Println(outputToStore)

		/*fmt.Printf("Earning Before Tax: %.1f\n", ebt)
		fmt.Printf("Profit: %.1f\n", profit)
		fmt.Printf("Ratio: %.1f\n", ratio)*/

		storeOutput(outputToStore)
	}
}

func validateInput(inputValue float64) (bool, error) {
	if inputValue <= 0 {
		err := errors.New("Invalid input")
		return false, err
	} else {
		return true, nil
	}
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeOutput(outputToStore string) {
	os.WriteFile("output.txt", []byte(outputToStore), 0644)
}
