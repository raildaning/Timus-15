package main

import (
	"fmt"
)

func main() {
	var i int
	sum := 0
	fmt.Scan(&i)
	switch {
	case i == 0:
		fmt.Println(1)
	case i < 0:
		for a := 1; a > i-1; a-- {
			sum += a
		}
		fmt.Println(sum)
	case i > 0:
		for a := 1; a < i+1; a++ {
			sum += a
		}
		fmt.Println(sum)
	}
}
