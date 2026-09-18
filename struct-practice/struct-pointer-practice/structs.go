package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getValueByPrompt("Enter your first name: ")
	userLastName := getValueByPrompt("Enter your last name: ")
	userBirthDate := getValueByPrompt("Enter your birthdate(MM/DD/YYYY): ")

	var appUser *user

	appUser = &user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	outputValues(appUser)
	clearUserName(appUser)
	outputValues(appUser)
}

func outputValues(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func clearUserName(u *user) {
	u.firstName = ""
	u.lastName = ""
}

func getValueByPrompt(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}
