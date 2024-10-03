package main

import "fmt"

func main() {

	fmt.Println("Bank App v0.1")
	balance := 0

	for {

		showMenu()

		var choice int
		_, _ = fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Print("Deposit amount: ")
			var amount int
			_, _ = fmt.Scanln(&amount)
			balance += amount
			fmt.Println("Deposited", amount)
			fmt.Println("Balance: ", balance)
		}
		if choice == 2 {
			fmt.Print("Withdraw amount: ")
			var amount int
			_, _ = fmt.Scanln(&amount)
			balance -= amount
			fmt.Println("Withdrew", amount)
			fmt.Println("Balance: ", balance)
		}
		if choice == 3 {
			fmt.Println("Balance: ", balance)
		}
		if choice == 4 {
			break
		} else {
			fmt.Println("Invalid choice")
		}

		fmt.Println("You chose", choice)

	}

	fmt.Println("Goodbye!")
}

func showMenu() {
	fmt.Println("What transaction would you like to do?")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Balance")
	fmt.Println("4. Exit")
	fmt.Println("Enter your choice: ")
}
