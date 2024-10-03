package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Bank App v0.1")
	balance := 0

	for {
		showMenu()

		choice := scanInt()

		switch choice {
		case 1:
			balance = deposit(balance)
		case 2:
			balance = withdraw(balance)
		case 3:
			fmt.Println("Balance: ", balance)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func withdraw(balance int) int {
	fmt.Print("Withdraw amount: ")
	amount := scanInt()

	if !validateAmount(amount) {
		return balance
	}

	newBalance := balance - amount

	if newBalance < 0 {
		fmt.Println("Insufficient funds")
		return balance
	}

	fmt.Println("Withdrew", amount)
	fmt.Println("Balance: ", newBalance)
	return newBalance
}

func deposit(balance int) int {
	fmt.Print("Deposit amount: ")
	amount := scanInt()

	if !validateAmount(amount) {
		return balance
	}

	balance += amount
	fmt.Println("Deposited", amount)
	fmt.Println("Balance: ", balance)
	return balance
}

func showMenu() {
	fmt.Println("What transaction would you like to do?")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Balance")
	fmt.Println("4. Exit")
	fmt.Println("Enter your choice: ")
}

func scanInt() int {
	const maxInputSize = 10
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, maxInputSize)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input := strings.TrimSpace(string(buf[:n]))

		amount, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a whole number.")
			break
		}
		return amount
	}

	return 0
}

func validateAmount(amount int) bool {
	if amount < 1 {
		fmt.Println("Enter an whole number greater than zero.")
		return false
	}

	return true
}
