package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdOn time.Time
}

func (u user) printUser() {
	fmt.Println("First Name:", strings.ToUpper(u.firstName))
	fmt.Println("Last Name:", strings.ToUpper(u.lastName))
	fmt.Println("Birth Date:", u.birthDate)
	fmt.Printf("Created On: %s\n", u.createdOn.Format(time.RFC3339))

	age, err := u.getAge()

	if err != nil {
		fmt.Println("Could not get age:", err)
		return
	}

	fmt.Println("Age:", age)
}

func (u user) getAge() (int, error) {
	birthDate, err := time.Parse(time.DateOnly, u.birthDate)

	if err != nil {
		return 0, errors.New("invalid birth date")
	}
	years := u.createdOn.Year() - birthDate.Year()

	if u.createdOn.Before(
		time.Date(
			u.createdOn.Year(),
			birthDate.Month(),
			birthDate.Day(),
			0,
			0,
			0,
			0,
			u.createdOn.Location())) {
		years--
	}

	return years, nil
}

const (
	PromptFirstName = "Enter your first name: "
	PromptLastName  = "Enter your last name: "
	PromptBirthDate = "Enter your birth date (YYYY-MM-DD): "
)

func main() {
	fn := promptAndGetData(PromptFirstName)
	ln := promptAndGetData(PromptLastName)
	dob := promptAndGetData(PromptBirthDate)

	var appUser = user{
		firstName: fn,
		lastName:  ln,
		birthDate: dob,
		createdOn: time.Now(),
	}

	appUser.printUser()
}

func promptAndGetData(prompt string) string {
	fmt.Print(prompt)
	var data string
	_, _ = fmt.Scanln(&data)
	return data
}
