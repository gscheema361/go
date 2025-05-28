package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter The Revenue: ")
	revenue := getUserInput("Enter The Revenue: ")
	expenses := getUserInput("Enter The Expenses: ")
	taxRate := getUserInput("Enter The Tax Rate: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Enter The Expenses: ")
	// getUserInput("Enter The Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Enter The tax rate: ")
	// getUserInput("Enter The tax rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinances(revenue, expenses, taxRate)

	// ebt := revenue - expenses

	// profit := ebt * (1 - taxRate/100)

	// ratio := ebt / profit

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinances(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
