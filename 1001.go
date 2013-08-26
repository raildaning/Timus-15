package main

import (
	"fmt"
	"math"
)

func main() {
	list := make([]float64, 0)
	var i float64

	for {
		k, _ := fmt.Scan(&i)

		if k != 1 {
			break
		}

		list = append(list, i)
	}
	
	for i := len(list) - 1; i >= 0; i-- {
		fmt.Printf("%.4f\n", math.Sqrt(list[i]))
	}
}
