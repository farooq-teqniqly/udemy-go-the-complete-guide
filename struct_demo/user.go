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
	birthDate time.Time
	createdOn time.Time
	age       int
}

func NewUser(firstName string, lastName string, birthDate string) (User, error) {
	fn := strings.TrimSpace(firstName)

	if fn == "" {
		return User{}, errors.New("first name cannot be empty")
	}

	ln := strings.TrimSpace(lastName)

	if ln == "" {
		return User{}, errors.New("last name cannot be empty")
	}

	parsedBirthDate, err := time.Parse(time.DateOnly, birthDate)

	if err != nil {
		return User{}, errors.New("invalid birth date")
	}

	now := time.Now()
	age := calculateAge(parsedBirthDate, now)

	return User{
		firstName: fn,
		lastName:  ln,
		birthDate: parsedBirthDate,
		createdOn: now,
		age:       age,
	}, nil
}

func (u User) print() {
	fmt.Println("First Name:", strings.ToUpper(u.firstName))
	fmt.Println("Last Name:", strings.ToUpper(u.lastName))
	fmt.Println("Birth date:", u.createdOn.Format(time.RFC3339))
	fmt.Println("Created On:", u.createdOn.Format(time.RFC3339))
	fmt.Println("Age:", u.age)
}

func calculateAge(birthDate time.Time, now time.Time) int {
	years := now.Year() - birthDate.Year()

	if now.Before(
		time.Date(
			now.Year(),
			birthDate.Month(),
			birthDate.Day(),
			0,
			0,
			0,
			0,
			now.Location())) {
		years--
	}

	return years
}
