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
	balance -= amount
	fmt.Println("Withdrew", amount)
	fmt.Println("Balance: ", balance)
	return balance
}

func deposit(balance int) int {
	fmt.Print("Deposit amount: ")
	amount := scanInt()
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
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		amount, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a whole number.")
			continue
		}
		return amount
	}
}
