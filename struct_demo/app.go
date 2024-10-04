package main

import (
	"fmt"
)

const (
	PromptFirstName = "Enter your first name: "
	PromptLastName  = "Enter your last name: "
	PromptBirthDate = "Enter your birth date (YYYY-MM-DD): "
)

func main() {
	fn := promptAndGetData(PromptFirstName)
	ln := promptAndGetData(PromptLastName)
	dob := promptAndGetData(PromptBirthDate)

	var appUser, err = NewUser(fn, ln, dob)

	if err != nil {
		fmt.Println("Could not create user.", err)
		return
	}

	appUser.print()
}

func promptAndGetData(prompt string) string {
	fmt.Print(prompt)
	var data string
	_, _ = fmt.Scanln(&data)
	return data
}
