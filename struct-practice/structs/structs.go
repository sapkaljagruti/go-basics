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

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	outputValues(appUser)
}

func outputValues(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getValueByPrompt(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}
