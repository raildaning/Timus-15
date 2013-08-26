package main

import (
	"fmt"
)

func main() {
	var i int
	
	fmt.Scan(&i)

	switch {
	case i <= 0:
	case i < 5:
		fmt.Println("few")
	case i < 10:
		fmt.Println("several")
	case i < 20:
		fmt.Println("pack")
	case i < 50:
		fmt.Println("lots")
	case i < 100:
		fmt.Println("horde")
	case i < 250:
		fmt.Println("throng")
	case i < 500:
		fmt.Println("swarm")
	case i < 1000:
		fmt.Println("zounds")
	default:
		fmt.Println("legion")
	}
}
