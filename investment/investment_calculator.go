package main

import (
	"fmt"
	"math"
)

const inflatinRate = 2.5

func main() {
	// const inflatinRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	// fmt.Print("Enter The Investment Amount: ")
	outputText("Enter The Investment Amount: ")
	fmt.Scan(&investmentAmount)
	// fmt.Print("Enter the Expected Return Rate: ")
	outputText("Enter the Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	// fmt.Print("Enter the Years: ")
	outputText("Enter the Years: ")
	fmt.Scan(&years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflatinRate/100, years)
	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value after inflation: %.1f\n", futureRealValue)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %v\n Future Value after inflation: %v\n", futureValue, futureRealValue)
	// fmt.Println("Future Real Value: ", futureRealValue)
	// fmt.Printf("Future Value: %.0f\n Future Value after inflation: %.1f\n", futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
	// fmt.Print(formattedRFV)
	// fmt.Printf(`Future Value: %v\n

	// Future Value after inflation: %v\n`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflatinRate/100, years)
	return fv, rfv
	// return
}
