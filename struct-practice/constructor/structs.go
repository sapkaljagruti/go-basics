package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required.")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func (u *user) outputValues() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getValueByPrompt("Enter your first name: ")
	userLastName := getValueByPrompt("Enter your last name: ")
	userBirthDate := getValueByPrompt("Enter your birthdate(MM/DD/YYYY): ")

	var appUser *user

	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputValues()
	appUser.clearUserName()
	appUser.outputValues()
}

func getValueByPrompt(prompt string) string {
	var value string
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
