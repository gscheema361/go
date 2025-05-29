package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceDataFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Fail to find balance File")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to Parse stored Value!")
	}
	return balance, nil
}
func writeDataToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceDataFromFile()
	if err != nil {
		fmt.Println("ERROR: ")
		fmt.Println(err)
		fmt.Println("------------------")
		panic("Can't Continue ,Sorry!")
	}
	fmt.Println("Welcome to go Bank!")

	for {
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
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				return
			}
			accountBalance += depositAmount
			writeDataToFile(accountBalance)
			fmt.Println("Your new balance is: ", accountBalance)

		} else if choice == 3 {
			fmt.Print("Your Withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= accountBalance {
				accountBalance -= withdrawalAmount
				fmt.Println("Your new balance is: ", accountBalance)
			} else {
				fmt.Println("Insufficient balance")
			}
		} else {
			fmt.Println("Exiting the program")
			//return
			break
		}

		fmt.Println("Your Choice:", choice)
	}
	fmt.Println("Thanks for choosing our bank.")
}
