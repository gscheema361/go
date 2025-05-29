package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string
	productNames = [4]string{"Product 1", "Product 2", "Product 3", "product 4"}
	prices := [4]float64{10.99, 9.99, 12.99, 11.99}
	fmt.Println(prices)
	fmt.Println(productNames[2])
	// titles  := [4]string {"Product A", "Product B", "Product C", "Product D"}
	fmt.Println(prices[3])
}
