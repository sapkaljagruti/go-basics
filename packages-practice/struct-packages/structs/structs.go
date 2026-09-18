package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getValueByPrompt("Enter your first name: ")
	userLastName := getValueByPrompt("Enter your last name: ")
	// userBirthDate := getValueByPrompt("Enter your birthdate(MM/DD/YYYY): ")

	var appUser *user.User

	appUser = &user.User{
		FirstName: userFirstName,
		LastName:  userLastName,
		// birthDate: userBirthDate, //it is not exported from user package, so we cannnot access this property here
		// createdAt: time.Now(), //it is not exported from user package, so we cannnot access this property here
	}

	outputValues(appUser)
	clearUserName(appUser)
	outputValues(appUser)
}

func outputValues(u *user.User) {
	fmt.Println(u.FirstName, u.LastName)
}

func clearUserName(u *user.User) {
	u.FirstName = ""
	u.LastName = ""
}

func getValueByPrompt(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}
