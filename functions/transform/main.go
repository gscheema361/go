package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 2, 1}
	doubled := transformNumbers(&numbers, double)
	trippled := transformNumbers(&numbers, tripple)
	fmt.Println(doubled)
	fmt.Println(trippled)
	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, multiplyNumber(val, 4))
	}
	return dNumbers
}

func transformNumbers(numbers *[]int, trasform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, trasform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return tripple
	}

}

func double(number int) int {
	return number * 2
}

func tripple(number int) int {
	return number * 3
}

func multiplyNumber(number, val int) int {
	return number * val
}
