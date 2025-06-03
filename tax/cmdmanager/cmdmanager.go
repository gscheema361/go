package cmdmanager

import "fmt"

type CMDmanager struct {
}

func (fm CMDmanager) ReadLines() ([]string, error) {
	fmt.Println("Please Enter Your prices.Confirm every price with Enter.")
	var prices []string
	for {
		var price string
		fmt.Print("Price:")
		fmt.Scan(&price)
		prices = append(prices, price)
		if price == "0" {
			break
		}
	}
	return prices, nil
}

func (fm CMDmanager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDmanager {
	return CMDmanager{}
}
