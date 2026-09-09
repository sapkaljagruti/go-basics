package main

import (
	"fmt"

	"example.com/bank-app/utility"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Print("\n\n\n\n\n\n")
	fmt.Println("Welcome to our bank!")
	fmt.Println("Reach us 24/7 ", randomdata.PhoneNumber())

	accountBalance, err := utility.GetFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Print("\n\n============================\n\n")
		fmt.Print("Error: ", err)
		fmt.Print("\n\n============================\n\n")
		//return
	}

	for {
		var choice int

		fmt.Print("\n\n============================\n\n")

		showOptions()

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

			utility.WriteFloatToFile(accountBalanceFile, accountBalance)

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

			utility.WriteFloatToFile(accountBalanceFile, accountBalance)

			fmt.Println("Your balance is updated. New amount is: ", accountBalance)
		default:
			fmt.Println("==========Goodbye!==========")
			utility.SayBye()
			fmt.Print("============================\n\n\n")
			return
		}
	}
}
