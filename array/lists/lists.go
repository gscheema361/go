package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 76.98}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string
// 	productNames = [4]string{"Product 1", "Product 2", "Product 3", "product 4"}
// 	prices := [5]float64{10.99, 9.99, 12.99, 11.99, 12.78}
// 	fmt.Println(prices)
// 	fmt.Println(productNames[2])
// 	// titles  := [4]string {"Product A", "Product B", "Product C", "Product D"}
// 	fmt.Println(prices[3])

// 	featuredPrices := prices[1:]
// 	fmt.Println(featuredPrices)
// 	highlighrdPrices := featuredPrices[:1]

// 	fmt.Println(highlighrdPrices)
// }
