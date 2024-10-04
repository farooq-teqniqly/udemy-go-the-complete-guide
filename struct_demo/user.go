package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdOn time.Time
}

func (u User) printUser() {
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

func (u User) getAge() (int, error) {
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
