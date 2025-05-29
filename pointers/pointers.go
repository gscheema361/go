package main

import "fmt"

func main() {
	age := 32
	// var agePointer *int

	agePointer := &age
	fmt.Println("Value of Address of age:", *agePointer) // Dereferencing the pointerto get the value
	// fmt.Println("Address of age:", agePointer)           // To get the memory address of the variable
	// fmt.Println("Age: ", age)

	// adultYears := getAdultYears(age)
	// fmt.Println("Adult years: ", adultYears)
	// adultYears := getAdultYears(agePointer)
	getAdultYears(agePointer)
	fmt.Println("Adult years: ", age)
}

// func getAdultYears(age int) int {
// 	return age - 18
// }

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
