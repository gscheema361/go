package main

import (
	"fmt"
)

type Product struct {
	Title string
	ID    string
	Price float64
}

func main() {
	// 1) Create an array of 3 hobbies
	hobbies := [3]string{"Reading", "Cycling", "Gaming"}
	fmt.Println("Hobbies:", hobbies)

	// 2) Output more data about the array
	fmt.Println("First hobby:", hobbies[0])
	secondAndThird := []string{hobbies[1], hobbies[2]}
	fmt.Println("Second and Third hobbies as list:", secondAndThird)

	// 3) Create two slices from the array (first and second elements)
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println("Slice1 (0:2):", slice1)
	fmt.Println("Slice2 (:2):", slice2)

	// 4) Re-slice to contain second and last element
	reSliced := []string{hobbies[1], hobbies[2]}
	fmt.Println("Re-sliced (second and last):", reSliced)

	// 5) Dynamic array (slice) for course goals
	courseGoals := []string{"Learn Go basics", "Build a real-world project"}
	fmt.Println("Course Goals:", courseGoals)

	// 6) Change second goal and add a third goal
	courseGoals[1] = "Master Go concurrency"
	courseGoals = append(courseGoals, "Contribute to open-source in Go")
	fmt.Println("Updated Course Goals:", courseGoals)

	// 7) Bonus: Product struct and dynamic list
	products := []Product{
		{Title: "Laptop", ID: "P1", Price: 1200.50},
		{Title: "Smartphone", ID: "P2", Price: 800.00},
	}
	fmt.Println("Products:", products)

	// Add a third product
	products = append(products, Product{Title: "Headphones", ID: "P3", Price: 150.75})
	fmt.Println("Updated Products:", products)
}
