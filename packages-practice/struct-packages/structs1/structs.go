package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getValueByPrompt("Enter your first name: ")
	userLastName := getValueByPrompt("Enter your last name: ")
	userBirthDate := getValueByPrompt("Enter your birthdate(MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputValues()
	appUser.ClearUserName()
	appUser.OutputValues()
}

func getValueByPrompt(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
