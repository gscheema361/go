package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your Balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your Deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your new balance is: ", accountBalance)
	}

	fmt.Println("Your Choice:", choice)
}
