package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) Output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Bob")
	userNames = append(userNames, "Alices")
	userNames[0] = "Alice"
	userNames[1] = "Charlie"

	fmt.Println(userNames)

	coursesRatings := make(floatMap, 4)
	coursesRatings["Math"] = 4.5
	coursesRatings["Science"] = 4.8
	coursesRatings["History"] = 4.2
	coursesRatings["English"] = 4.9
	coursesRatings["Geography"] = 4.7
	coursesRatings["Art"] = 4.3
	coursesRatings["Music"] = 4.6
	coursesRatings["Dance"] = 4.4
	coursesRatings.Output()

}
