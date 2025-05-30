package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	double := createTransformer(2)
	tripple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	trippled := transformNumbers(&numbers, tripple)
	fmt.Println("Original:", numbers)
	fmt.Println("Transformed:", transformed)
	fmt.Println("Doubled:", doubled)
	fmt.Println("Tripled:", trippled)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
