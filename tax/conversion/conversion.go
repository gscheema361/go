package conversion

import (
	"fmt"
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			fmt.Println("Converting price to float failed!")
			fmt.Println(err)
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
