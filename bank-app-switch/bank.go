package main

import "fmt"

func main() {
	fmt.Print("\n\n\n\n\n\n")
	fmt.Println("Welcome to our bank!")

	var accountBalance = 1000.0

	for {
		var choice int

		fmt.Print("\n\n============================\n\n")
		fmt.Println("What would you like to do?")
		fmt.Println("1. View Balance:")
		fmt.Println("2. Deposit Money:")
		fmt.Println("3. Withdraw Money:")
		fmt.Println("4. Exit:")

		fmt.Print("\n\nPlease enter your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("\n\nYour balance is ", accountBalance)
		case 2:
			var depositAmount float64

			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("The entered amount to depost is not valid. Please try again.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your balance is updated. New amount is: ", accountBalance)
		case 3:
			var withdrawAmount float64

			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("The entered amount to withdraw is not valid. Please try again.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You cannot withdraw more than you have. Please try again.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your balance is updated. New amount is: ", accountBalance)
		default:
			fmt.Println("==========Goodbye!==========")
			fmt.Println("Thank you for choosing our bank.")
			fmt.Print("============================\n\n\n")
			return
		}
	}
}
