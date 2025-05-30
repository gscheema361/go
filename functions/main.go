package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 10, 15, -5)
	anotherSun := sumup(1, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSun)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum + startingValue
}
